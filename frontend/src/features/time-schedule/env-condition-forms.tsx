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
import { ClimateData, Condition } from "@/types/api";
import { ChangeEvent, useEffect, useState } from "react";

interface EnvConditionFormsProps {
  initialCondition: Condition | null;
  onFormsChange: (updatedData: Condition) => void;
}

export const EnvConditionForms = (props: EnvConditionFormsProps) => {
  const { initialCondition, onFormsChange } = props;
  const [timeScheduleInfo] = useTimeScheduleInfo();
  const [selectedClimateData, setSelectedClimateData] = useState<
    ClimateData | undefined
  >(undefined);
  const [selectedCmpOpe, setSelectedCmpOpe] = useState<string>("");
  const [setPoint, setSetPoint] = useState<string>("");
  const [condition, setCondition] = useState<Condition | null>(
    initialCondition
  );

  useEffect(() => {
    if (!initialCondition) {
      return;
    }

    const selectedClimateData: ClimateData | undefined =
      timeScheduleInfo.climateData.find(
        (climateData) =>
          climateData.id === initialCondition.selected_climate_data_id
      );

    if (selectedClimateData) {
      setSelectedClimateData(selectedClimateData);
    }

    setSelectedCmpOpe(String(initialCondition.selected_comparison_operator_id));
    setSetPoint(String(initialCondition.set_point));
  }, []);

  const handleClimateDataChange = (event: SelectChangeEvent) => {
    const climateDataID = parseInt(event.target.value);
    const climateDataRec = timeScheduleInfo?.climateData.find(
      (data) => data.id === climateDataID
    );

    if (!condition) {
      return;
    }

    const newCondition: Condition = {
      ...condition,
      selected_climate_data_id: climateDataID,
    };

    setSelectedClimateData(climateDataRec);
    setCondition(newCondition);
    onFormsChange(newCondition);
  };

  const handleCmpOpeChange = (event: SelectChangeEvent) => {
    const newCmpID: string = event.target.value;
    const intNewCmpID: number = parseInt(newCmpID);

    if (!condition) {
      return;
    }

    const newCondition: Condition = {
      ...condition,
      selected_comparison_operator_id: intNewCmpID,
    };
    setSelectedCmpOpe(newCmpID);
    setCondition(newCondition);
    onFormsChange(newCondition);
  };

  const handleSetPointChange = (
    event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const newValue: string = event.target.value;
    const newSetPoint: number = parseInt(newValue);

    if (!condition) {
      return;
    }

    const newCondition: Condition = {
      ...condition,
      set_point: newSetPoint,
    };

    setSetPoint(newValue);
    setSetPoint(newValue === "" ? "" : newValue);
    setCondition(newCondition);
    onFormsChange(newCondition);
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
            value={selectedCmpOpe != "0" ? selectedCmpOpe : ""}
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
          value={setPoint}
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
          onChange={handleSetPointChange}
        />
      </Box>
    </Box>
  );
};
