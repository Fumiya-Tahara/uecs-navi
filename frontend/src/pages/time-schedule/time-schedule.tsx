import { Box, Tab, Tabs, CircularProgress } from "@mui/material";
import { TimeTable } from "./time-table";
import { Navigation } from "@/layouts/navigation";
import { useState, useEffect } from "react";
import { ClimateData, TimeScheduleResponse } from "@/types/api";
import {
  getClimateDatas,
  getTimeSchedules,
} from "@/features/api/mocks/setting-time-schedule-api";
import { useTimeScheduleInfo } from "@/hooks/time-schedule-info-context";
import { useM304IDs } from "@/hooks/m304ids-context";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function a11yProps(index: number) {
  return {
    id: `setting-device-tabpanel-${index}`,
    "aria-controls": `setting-device-tabpanel-${index}`,
  };
}

function TimeScheduleTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <Box
      role="tabpanel"
      hidden={value !== index}
      id={`time-schedule-tabpanel-${index}`}
      aria-labelledby={`time-schedule-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
    </Box>
  );
}

export function TimeSchedule() {
  const [m304IDs] = useM304IDs();
  const [selectedTab, setSelectedTab] = useState<number>(0);
  const [loading, setLoading] = useState<boolean>(true);

  const [timeScheduleInfo, setTimeScheduleInfo] = useTimeScheduleInfo();

  useEffect(() => {
    const fetchInitData = async () => {
      setLoading(true);
      const climateData: ClimateData[] = getClimateDatas();
      const timeScheduleMap: Map<number, TimeScheduleResponse> = new Map();

      if (!m304IDs.m304IDs) {
        setLoading(false);
        return;
      }

      for (const m304ID of m304IDs.m304IDs) {
        const timeScheduleRes: TimeScheduleResponse | null =
          await getTimeSchedules(m304ID);
        if (timeScheduleRes) {
          timeScheduleMap.set(m304ID, timeScheduleRes);
        }
      }

      setTimeScheduleInfo({
        climateData: climateData,
        m304TimeScheduleMap: timeScheduleMap,
      });
      setLoading(false);
    };

    fetchInitData();
  }, [m304IDs, setTimeScheduleInfo]);

  const handleTabChange = (_event: React.SyntheticEvent, newValue: number) => {
    console.log(newValue);
    setSelectedTab(newValue);
  };

  return (
    <Navigation>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={selectedTab}
          onChange={handleTabChange}
          aria-label="basic tabs example"
        >
          {m304IDs.m304IDs?.map((m304ID, index) => (
            <Tab key={index} label={m304ID} {...a11yProps(index)} />
          ))}
        </Tabs>
      </Box>
      {loading ? (
        <Box sx={{ display: "flex", justifyContent: "center", mt: 2 }}>
          <CircularProgress />
        </Box>
      ) : (
        m304IDs.m304IDs?.map((m304ID, index) => (
          <TimeScheduleTabPanel key={index} index={index} value={selectedTab}>
            <TimeTable
              initialSchedules={
                timeScheduleInfo?.m304TimeScheduleMap.get(m304ID) || null
              }
              workflows={
                timeScheduleInfo?.m304TimeScheduleMap.get(m304ID)?.workflows ||
                []
              }
            />
          </TimeScheduleTabPanel>
        ))
      )}
    </Navigation>
  );
}
