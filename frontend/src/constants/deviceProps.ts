export type DeviceProps = {
  id: number;
  device_id: number;
  display_name: string;
  active_status: string;
  latest_device_point: {
    lat: number;
    lng: number;
    device_state: {
      drive_status: string;
    };
  };
  Visible: boolean;
  IconUrl: string;
};
