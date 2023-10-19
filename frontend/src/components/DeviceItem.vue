<template>
  <div class="device-item card" @click="handleClick">
    <div :class="isSelectedClass">
      <div class="card-body">
        <div class="row">
          <div class="col-md-2 image-container">
            <div v-if="device.IconUrl" class="image-center">
              <img
                class="circular-image"
                :src="require(`@/assets/${device.IconUrl}`)"
                alt="Device Image"
              />
            </div>
            <div v-else class="colored-circle image-center">
              {{ device.display_name.charAt(0) }}
            </div>
          </div>
          <div class="col-md-10">
            <h5 class="device-name card-title text-center">
              {{ device.display_name }}
            </h5>
            <div class="device-info mb-0 mt-3">
              <div class="row">
                <div class="drive-status col-md-6">
                  <span
                    v-show="
                      device.latest_device_point.device_state.drive_status ==
                      'driving'
                    "
                  >
                    <i class="fa-solid fa-power-off" style="color: #00f504"></i>
                    Driving
                  </span>
                  <span
                    v-show="
                      device.latest_device_point.device_state.drive_status ==
                      'idle'
                    "
                  >
                    <i class="fa-solid fa-power-off" style="color: #fbb904"></i>
                    Idle
                  </span>
                  <span
                    v-show="
                      device.latest_device_point.device_state.drive_status ==
                      'off'
                    "
                  >
                    <i class="fa-solid fa-power-off"></i>
                    Ignition off
                  </span>
                  <p class="card-text">Status: {{ device.active_state }}</p>
                </div>
                <div class="location col-md-6 text-right">
                  <p class="card-text">
                    Lat: {{ device.latest_device_point.lat.toFixed(4) }}
                  </p>
                  <p class="card-text">
                    Long: {{ device.latest_device_point.lng.toFixed(4) }}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding-left: 0;
  padding-right: 0;
}

.image-center {
  display: inline-block;
}

.circular-image {
  border-radius: 50%;
  width: 50px;
  height: 50px;
  object-fit: cover;
  margin: auto;
}

.colored-circle {
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  width: 50px;
  height: 50px;
  max-width: 100%;
  max-height: 100%;
  background-color: #3498db;
  color: white;
  margin: auto;
  font-size: 20px;
}

.device-name {
  margin: 0px;
}

.isSelected {
  background-color: #95cff4;
}

.location {
  margin: 0px;
}

.drive_status {
  float: left;
  margin: 0px;
}
</style>

<script>
import { DeviceProps } from "../constants/deviceProps";

export default {
  props: {
    device: {
      type: DeviceProps,
      required: true,
    },
    selectedDeviceId: {
      type: String,
      required: true,
    },
  },
  emits: ["device-selected"],
  methods: {
    handleClick() {
      this.$emit("device-selected", this.device);
    },
  },
  computed: {
    isSelectedClass() {
      return {
        isSelected: this.device.device_id === this.selectedDeviceId,
        "": !this.device.device_id === this.selectedDeviceId,
      };
    },
  },
};
</script>
