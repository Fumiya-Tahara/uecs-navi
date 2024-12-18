import { ClimateDataResponse } from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type timeScheduleInfoState = {
  climate_data: ClimateDataResponse[];
};

type timeScheduleInfoContextType = [
  timeScheduleInfoState,
  Dispatch<SetStateAction<timeScheduleInfoState>>
];

const timeScheduleInfoContext = createContext<timeScheduleInfoContextType>([
  {
    climate_data: [],
  },
  () => {},
]);

export const TimeScheduleInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [timeScheduleInfo, setTimeScheduleInfo] =
    useState<timeScheduleInfoState>({
      climate_data: [],
    });

  return (
    <timeScheduleInfoContext.Provider
      value={[timeScheduleInfo, setTimeScheduleInfo]}
    >
      {children}
    </timeScheduleInfoContext.Provider>
  );
};

export const useTimeScheduleInfo = () => {
  return useContext(timeScheduleInfoContext);
};

export default timeScheduleInfoContext;
