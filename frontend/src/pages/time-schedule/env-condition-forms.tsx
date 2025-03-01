import { useTimeScheduleInfo } from "@/hooks/time-schedule-info-context";
import {
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  InputAdornment,
  SelectChangeEvent,
} from "@mui/material";
import { ClimateData } from "@/types/api";
import { useState } from "react";

export const EnvConditionForms = () => {
  const [timeScheduleInfo] = useTimeScheduleInfo();
  const [selectedClimateData, setSelectedClimateData] = useState<
    ClimateData | undefined
  >(undefined);
  const [selectedCmpOpe, setSelectedCmpOpe] = useState<string>("");

  const handleClimateDataChange = (event: SelectChangeEvent) => {
    const climateDataID = parseInt(event.target.value);
    const climateDataRec = timeScheduleInfo?.climateData.find(
      (data) => data.id === climateDataID
    );

    setSelectedClimateData(climateDataRec);
  };

  const handleCmpOpeChange = (event: SelectChangeEvent) => {
    setSelectedCmpOpe(event.target.value as string);
  };

  return (
    <Box>
      <Box sx={{ display: "flex" }}>
        <FormControl sx={{ flex: 3 }}>
          <InputLabel shrink id={`climate-data-node-select-label`} size="small">
            気象データ
          </InputLabel>
          <Select
            labelId={`climate-data-node-select-label`}
            id={`climate-data-node-select`}
            value={selectedClimateData ? String(selectedClimateData.id) : ""}
            label="気象データ"
            onChange={handleClimateDataChange}
            size="small"
            notched
          >
            {timeScheduleInfo?.climateData.map((data) => (
              <MenuItem key={data.id} value={data.id}>
                {data.name}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
        <FormControl sx={{ flex: 2 }}>
          <InputLabel shrink id="comp-ope-label" size="small">
            比較
          </InputLabel>
          <Select
            id="comp-ope"
            labelId="comp-ope-label"
            value={selectedCmpOpe}
            size="small"
            onChange={handleCmpOpeChange}
            sx={{ marginX: 1 }}
            label="比較"
            notched
          >
            <MenuItem value={1}>{"="}</MenuItem>
            <MenuItem value={2}>{">"}</MenuItem>
            <MenuItem value={3}>{"<"}</MenuItem>
            <MenuItem value={4}>{"≥"}</MenuItem>
            <MenuItem value={5}>{"≤"}</MenuItem>
          </Select>
        </FormControl>
        <TextField
          type="number"
          size="small"
          slotProps={{
            input: {
              endAdornment: (
                <InputAdornment position="end">
                  {selectedClimateData?.unit}
                </InputAdornment>
              ),
            },
          }}
          inputProps={{ step: "0.1" }}
          sx={{ marginRight: "8px", flex: 3 }}
        />
      </Box>
    </Box>
  );
};
