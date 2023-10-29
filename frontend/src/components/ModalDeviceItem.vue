<template>
  <div class="device-item">
    <div class="image-container" @click="openImageUpload">
      <input
        type="file"
        accept="image/*"
        style="display: none"
        ref="imageInput"
        @change="handleImageUpload"
      />
      <div v-if="device.IconUrl" class="image-center">
        <div v-if="device.IconUrl.charAt(0) == 'i'">
          <img
            class="profile-picture"
            :src="require(`@/assets/${device.IconUrl}`)"
            alt="Device Image"
          />
        </div>
        <div v-else>
          <img
            class="profile-picture"
            :src="device.IconUrl"
            alt="Device Image"
          />
        </div>
      </div>
      <div v-else class="colored-circle image-center">
        <i class="fa-solid fa-plus"></i>
      </div>
    </div>
    <div class="device-info">
      <h5 class="device-name">{{ device.display_name }}</h5>
    </div>
    <div class="visibility-icon" @click="toggleVisibility">
      <i
        class="fas"
        :class="{ 'fa-eye': !device.Visible, 'fa-eye-slash': device.Visible }"
      >
      </i>
    </div>
  </div>
</template>

<style scoped>
.device-item {
  display: flex;
  align-items: center;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  margin: 10px;
}

.profile-picture {
  flex: 0 0 50px;
  border-radius: 50%;
  width: 50px;
  max-width: 50px;
  max-height: 50px;
  margin: auto;
  overflow: hidden;
  object-fit: cover;
  aspect-ratio: 1/1;
}

.profile-picture img {
  width: 100%;
  height: 100%;
}

.device-info {
  flex: 1;
  text-align: center;
}

.device-name {
  font-weight: bold;
  margin: 0;
}

.visibility-icon {
  flex: 0 0 20px;
  text-align: right;
  color: #007bff;
  cursor: pointer;
  font-size: 24px;
}

.image-center {
  display: inline-block;
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
  background-color: #5f6366;
  color: white;
  margin: auto;
  font-size: 20px;
}

.image-container {
  cursor: pointer;
}
</style>

<script>
export default {
  props: {
    device: {
      type: Object,
      required: true,
    },
  },
  emits: ["toggle-visibility", "image-uploaded"],
  methods: {
    toggleVisibility() {
      this.$emit("toggle-visibility", this.device);
    },
    openImageUpload() {
      this.$refs.imageInput.click();
    },
    handleImageUpload(event) {
      const file = event.target.files[0];
      if (file) {
        this.$emit("image-uploaded", { device: this.device, image: file });
      }
    },
  },
};
</script>
