<template>
  <div class="modal" tabindex="-1" role="dialog">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">Edit Preferences</h3>
          <button class="btn btn-link" type="button" @click="closeModal">
            <i class="fa fa-window-close" style="font-size: 28px; color: red">
            </i>
          </button>
        </div>

        <div v-if="error" class="error-banner">
          {{ error }}
          <button @click="clearError">Close</button>
        </div>
        <div class="modal-body">
          <div class="devices-section">
            <h4 class="device-header">Visible Devices</h4>
            <div v-if="modalVisibleDevices.length == 0">
              <p class="mb-3">No visible devices</p>
            </div>
            <div v-else>
              <div v-for="device in modalVisibleDevices" :key="device.id">
                <ModalDeviceItem
                  :device="device"
                  @toggle-visibility="toggleVisibility"
                  @image-uploaded="handleImageUpload"
                />
              </div>
            </div>
            <h4>Hidden Devices</h4>
            <div v-if="modalHiddenDevices.length == 0">
              <p>No hidden devices</p>
            </div>
            <div v-else>
              <div v-for="device in modalHiddenDevices" :key="device.id">
                <ModalDeviceItem
                  :device="device"
                  @toggle-visibility="toggleVisibility"
                  @image-uploaded="handleImageUpload"
                />
              </div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" @click="closeModal">
            Cancel
          </button>
          <button type="button" class="btn btn-primary" @click="saveChanges">
            Save Changes
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: rgba(0, 0, 0, 0.5);
  z-index: 9999;
}

.error-banner {
  background-color: #ff7066;
  color: white;
  padding: 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.error-banner button {
  background-color: transparent;
  border: none;
  color: white;
  font-weight: bold;
  cursor: pointer;
}

.modal-content {
  background-color: #fff;
  margin-right: 0;
  padding: 20px;
  border-radius: 4px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  max-width: 100%;
  max-height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.modal-dialog {
  background-color: #fff;
  padding: 20px;
  border-radius: 4px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  height: 100%;
  overflow: auto;
  max-height: 100%;
}

.modal-footer {
  flex: 0 0 auto;
  border: none;
}

.modal-body {
  flex: 1 1 auto;
  overflow-y: auto;
}

@media (min-width: 768px) {
  .modal-dialog {
    max-width: 50%;
    max-height: 80%;
  }
}

.devices-section {
  margin: 0;
}
</style>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { DeviceProps } from "../constants/deviceProps";
import ModalDeviceItem from "./ModalDeviceItem.vue";
import axios from "axios";

export default defineComponent({
  components: {
    ModalDeviceItem,
  },
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    visibleDevices: {
      type: Array as PropType<DeviceProps[]>,
      required: true,
    },
    hiddenDevices: {
      type: Array as PropType<DeviceProps[]>,
      required: true,
    },
  },
  emits: ["close", "update"],
  data() {
    return {
      modalVisibleDevices: [] as DeviceProps[],
      modalHiddenDevices: [] as DeviceProps[],
      selectedImages: {} as { [key: string]: File },
      updatedVisibility: {} as { [key: string]: boolean },
      error: false as string | false,
    };
  },
  mounted() {
    this.modalVisibleDevices = this.deepCopyList(this.visibleDevices);
    this.modalHiddenDevices = this.deepCopyList(this.hiddenDevices);
  },
  methods: {
    closeModal() {
      this.$emit("close");
    },
    clearError() {
      this.error = false;
    },
    async saveChanges() {
      try {
        const visibilityUpdates = [] as {
          deviceId: string;
          visible: boolean;
        }[];

        for (const deviceId in this.updatedVisibility) {
          visibilityUpdates.push({
            deviceId,
            visible: this.updatedVisibility[deviceId],
          });
        }

        const imageUploads = new FormData();

        for (const deviceId in this.selectedImages) {
          imageUploads.append(deviceId, this.selectedImages[deviceId]);
        }

        const visibilityUpdateUrl =
          "http://localhost:3000/api/devices/update/visibility";
        const imageUploadUrl = "http://localhost:3000/api/devices/update/icon";

        const [visibilityResponse, imageUploadResponse] = await Promise.all([
          axios.patch(visibilityUpdateUrl, visibilityUpdates, {
            headers: {
              "Content-Type": "application/json",
            },
          }),
          axios.patch(imageUploadUrl, imageUploads, {
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }),
        ]);

        if (visibilityResponse.status !== 200) {
          this.error =
            "Error updating device visibility. Please try again later.";
        } else if (imageUploadResponse.status !== 200) {
          this.error = "Error updating device icon. Please try again later.";
        } else {
          this.$emit("update");
          this.closeModal();
        }
      } catch (error) {
        this.error =
          "Error updating device preferences. Please try again later.";
      }
    },
    toggleVisibility(device: DeviceProps) {
      if (device.Visible) {
        const index = this.modalVisibleDevices.findIndex(
          (obj) => obj.device_id == device.device_id
        );
        const [movedDevice] = this.modalVisibleDevices.splice(index, 1);
        this.modalHiddenDevices.push(movedDevice);
        device.Visible = false;
        this.updatedVisibility[device.device_id] = false;
      } else {
        const index = this.modalHiddenDevices.findIndex(
          (obj) => obj.device_id == device.device_id
        );
        const [movedDevice] = this.modalHiddenDevices.splice(index, 1);
        this.modalVisibleDevices.push(movedDevice);
        device.Visible = true;
        this.updatedVisibility[device.device_id] = true;
      }
    },
    handleImageUpload({ device, image }: { device: DeviceProps; image: File }) {
      const reader = new FileReader();
      if (image) {
        reader.onload = () => {
          device.IconUrl = reader.result as string;
        };
        reader.readAsDataURL(image);
        this.selectedImages[device.device_id] = image;
      }
    },
    deepCopyList(list: DeviceProps[]) {
      return JSON.parse(JSON.stringify(list));
    },
  },
});
</script>
