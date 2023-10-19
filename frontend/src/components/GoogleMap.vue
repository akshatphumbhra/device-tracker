<template>
  <GMapMap
    :center="position"
    :zoom="15"
    map-type-id="terrain"
    style="width: 100%; height: 100vh"
  >
    <GMapMarker
      :position="position"
      :clickable="true"
      :draggable="true"
      @click="center = position"
      :icon="{
        url: 'https://img.icons8.com/color-glass/48/center-direction.png',
        scaledSize: { width: 45, height: 45 },
      }"
    />
  </GMapMap>
</template>

<script>
export default {
  props: {
    lat: {
      type: Number,
      required: true,
    },
    lng: {
      type: Number,
      required: true,
    },
  },
  computed: {
    position() {
      return {
        lat: this.lat,
        lng: this.lng,
      };
    },
  },
};
</script>
