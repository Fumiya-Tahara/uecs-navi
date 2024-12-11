import { useDnD } from "@/hooks/dnd-context";
import { DragEvent } from "react";
import { Box, Divider, Typography } from "@mui/material";
import PlayCircleIcon from "@mui/icons-material/PlayCircle";
import { useState, useEffect } from "react";
import { DeviceResponse } from "@/types/api";
import { getDevices } from "@/mocks/workflow-api";

export const Sidebar = () => {
  const [, setType] = useDnD();
  const [fetchedDevices, setFetchedDevices] = useState<DeviceResponse[]>([]);

  useEffect(() => {
    const fetchDevices = async () => {
      const devicesRes: DeviceResponse[] = await getDevices();
      setFetchedDevices(devicesRes);
    };

    fetchDevices();
  }, []);

  const onDragStart = (
    event: DragEvent,
    nodeType: string,
    nodeData: object
  ) => {
    setType(nodeType);

    event.dataTransfer.setData(
      "application/reactflow",
      JSON.stringify(nodeData)
    );
    event.dataTransfer.effectAllowed = "move";
  };

  return (
    <Box sx={{ height: "100%", width: "300px", backgroundColor: "#E0E0E0" }}>
      <Box sx={{ display: "flex", flexDirection: "column" }}>
        <Box
          sx={{
            mx: 4,
            my: 2,
            border: "1px solid #000",
            borderRadius: "10px",
            backgroundColor: "#FFF",
          }}
          onDragStart={(event) =>
            onDragStart(event, "operation", {
              devicesList: fetchedDevices,
            })
          }
          draggable
        >
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              gap: 1,
              borderRadius: "10px 10px 0 0",
              color: "#FFF",
              backgroundColor: "#66BB6A",
              padding: "4px 8px 4px 8px",
            }}
          >
            <PlayCircleIcon />
            <Typography variant="body1">Operation</Typography>
          </Box>
          <Divider />
          <Box sx={{ padding: 1, textAlign: "center" }}>操作</Box>
        </Box>
      </Box>
    </Box>
  );
};
