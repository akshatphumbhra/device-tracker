<template>
  <div class="home-container">
    <div class="device-list-container">
      <DeviceList @update-mapmarker="updateMap" />
    </div>
    <div class="map-container">
      <GoogleMap :lat="lat" :lng="lng" />
    </div>
  </div>
</template>

<style scoped>
.device-list-container {
  flex: 0 0 30%;
  height: 100vh;
}

.home-container {
  display: flex;
}

.map-container {
  flex: 1;
  height: 100vh;
}
</style>

<script lang="ts">
import DeviceList from "@/components/DeviceList.vue";
import GoogleMap from "@/components/GoogleMap.vue";
import { DeviceProps } from "../constants/deviceProps";
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    DeviceList,
    GoogleMap,
  },
  data() {
    return {
      lat: 0,
      lng: 0,
    } as {
      lat: number;
      lng: number;
    };
  },
  methods: {
    updateMap(device: DeviceProps) {
      this.lat = device.latest_device_point.lat;
      this.lng = device.latest_device_point.lng;
    },
  },
});
</script>
