<template>
  <div class="list-header">
    <div class="d-flex justify-content-between align-items-start">
      <h2 class="list-title ml-3">Devices</h2>
      <div class="settings">
        <button
          class="btn btn-link custom-link p-0 pt-1"
          @click="showModal = true"
        >
          <i class="fa fa-cog" aria-hidden="true"></i>
        </button>

        <PreferenceModal
          v-if="showModal"
          :visibleDevices="visibleDevices"
          :hiddenDevices="hiddenDevices"
          :visible="showModal"
          @close="showModal = false"
          @update="fetchDeviceData"
        />
        <button
          class="btn btn-link custom-link no-padding"
          @click="fetchDeviceData"
        >
          <i class="m-3 fa fa-refresh" aria-hidden="true"></i>
        </button>
      </div>
    </div>
    <hr class="m-0" />
    <div class="d-flex justify-content-between align-items-center">
      <button class="btn btn-link custom-link" @click="toggleSort">
        <span v-show="sort == 'asc'">
          <i class="fa fa-caret-up"></i>
        </span>
        <span v-show="sort == 'desc'">
          <i class="fa fa-caret-down"></i>
        </span>
        <span class="bolder m-1">SORT</span>
      </button>
      <p class="num-devices">
        {{ visibleDevices.length }} /
        {{ visibleDevices.length + hiddenDevices.length }}
      </p>
    </div>
  </div>
  <div v-if="isLoading && visibleDevices.length == 0">
    <LoadingSpinner />
  </div>
  <div v-else>
    <div class="device-list">
      <div
        class="device-item"
        v-for="device in visibleDevices"
        :key="device.id"
      >
        <DeviceItem
          :device="device"
          :selectedDeviceId="selectedDevice.device_id"
          @device-selected="updateMapmarker"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.custom-link {
  text-decoration: none; /* Remove underline */
  color: #333; /* Change link color to gray (#333) */
  float: left;
}

.bolder {
  font-weight: bolder;
}

.num-devices {
  float: right;
  margin-top: 0.5rem;
  margin-right: 0.5rem;
}

.list-title {
  margin-top: 0.6rem;
  margin-left: 0.5rem;
  float: left;
}

.settings {
  display: flex;
  align-items: center;
}

.no-padding {
  padding-left: 0;
  padding-right: 0;
  padding-bottom: 0;
}

.device-list {
  float: left;
  width: 100%;
  max-height: calc(100vh - 100px);
  overflow-y: scroll;
}

.device-item {
  cursor: pointer;
  transition: background-color 0.3s;
}

.device-item:hover {
  background-color: #f0f0f0;
}

.list-header {
  height: 100px;
}
</style>

<script lang="ts">
import axios from "axios";
import DeviceItem from "./DeviceItem.vue";
import PreferenceModal from "./PreferenceModal.vue";
import LoadingSpinner from "./LoadingSpinner.vue";
import { defineComponent, Ref, ref } from "vue";
import { DeviceProps } from "../constants/deviceProps";

export default defineComponent({
  components: {
    DeviceItem,
    LoadingSpinner,
    PreferenceModal,
  },
  emits: ["update-mapmarker"],
  data() {
    return {
      visibleDevices: [] as DeviceProps[],
      hiddenDevices: [] as DeviceProps[],
      selectedDevice: {} as DeviceProps,
      sort: "asc",
      showModal: false,
      isLoading: false,
    };
  },
  mounted() {
    this.fetchDeviceData();
    // Update data every minute
    setInterval(this.fetchDeviceData, 60000);
  },
  methods: {
    fetchDeviceData() {
      this.isLoading = true;
      axios
        .get("http://localhost:3000/api/devices/", {
          headers: {
            "Content-Type": "application/json",
          },
        })
        .then((response) => {
          this.visibleDevices = [];
          this.hiddenDevices = [];

          // Filter the data into visible and hidden devices
          response.data.forEach((device: DeviceProps) => {
            if (device.Visible) {
              this.visibleDevices.push(device);
            } else {
              this.hiddenDevices.push(device);
            }
          });
          const sortDirection = localStorage.getItem("sort");
          if (sortDirection) {
            this.sort = sortDirection;
          }
          this.sortList(this.sort);
          if (Object.keys(this.selectedDevice).length === 0) {
            this.selectedDevice = this.visibleDevices[0];
          }
          this.isLoading = false;
          this.updateMapmarker(this.selectedDevice);
        })
        .catch((error) => {
          console.error(error);
        });
    },
    updateMapmarker(device: DeviceProps) {
      this.selectedDevice = device;
      this.$emit("update-mapmarker", device);
    },
    sortList(sortDirection: string) {
      if (sortDirection == "asc") {
        this.visibleDevices.sort((a, b) =>
          a.display_name.localeCompare(b.display_name)
        );
      } else {
        this.visibleDevices.sort((a, b) =>
          b.display_name.localeCompare(a.display_name)
        );
      }
    },
    toggleSort() {
      this.sort = this.sort == "asc" ? "desc" : "asc";
      localStorage.setItem("sort", this.sort);
      this.sortList(this.sort);
    },
  },
});
</script>
