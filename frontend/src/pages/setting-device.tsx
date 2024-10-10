import { DeviceCard } from "../components/setting-device"
import { Navigation } from "../layouts/navigation"
import { Box, Tabs, Tab } from "@mui/material"
import Grid from "@mui/material/Grid2";
import { useState, useEffect } from "react";
import { HouseResponse, JoinedDeviceResponse } from "../types/api";
import { getDevices, getHouses } from "../mocks/setting_device_api";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function SettingDeviceTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`setting-device-tabpanel-${index}`}
      aria-labelledby={`setting-device-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `setting-device-tabpanel-${index}`,
    "aria-controls": `setting-device-tabpanel-${index}`,
  };
}

export default function SettingDevice() {
  const [houses, setHouses] = useState<HouseResponse[]>([]);
  const [devicesMap, setDevicesMap] = useState<Map<number, JoinedDeviceResponse[]>>(new Map());
  
  const [value, setValue] = useState(0);

  useEffect(() => {
    const housesRes: HouseResponse[] = getHouses();
    

    // 複数のハウスIDから同時にdeviceを取得するAPIを用意するべき
    const devicesMap: Map<number, JoinedDeviceResponse[]> = new Map();
    housesRes.forEach((house) => {
      const devicesRes: JoinedDeviceResponse[] = getDevices(house.id);
      devicesMap.set(house.id, devicesRes);
    })

    setHouses(housesRes);
    setDevicesMap(devicesMap);
  }, []);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
      <Navigation>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
          <Tabs value={value} onChange={handleChange} aria-label="basic tabs example">
            {houses.map((house, index) => (
              <Tab key={index} label={house.name} {...a11yProps(index)}/>
            ))}
          </Tabs>
        </Box>
          {houses.map((house, index) => (
            <SettingDeviceTabPanel key={index} index={index} value={value}>
              <Grid container spacing={4}>
                {devicesMap.get(house.id)?.map((device, deviceIndex) => (
                  <Grid size={3} key={deviceIndex}>
                    <DeviceCard 
                      name={device.name} 
                      setPoint={device.setPoint} 
                      duration={device.duration}
                      climateData={device.climateData}
                      unit={device.unit}
                    />
                  </Grid>
                ))}
              </Grid>
            </SettingDeviceTabPanel>
          ))}
      </Navigation>
  )
}
