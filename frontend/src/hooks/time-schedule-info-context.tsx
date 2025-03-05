import { ClimateData, TimeScheduleResponse } from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

export type timeScheduleInfoState = {
  climateData: ClimateData[];
  m304TimeScheduleMap: Map<number, TimeScheduleResponse>;
};

type timeScheduleInfoContextType = [
  timeScheduleInfoState,
  Dispatch<SetStateAction<timeScheduleInfoState>>
];

const timeScheduleInfoContext = createContext<timeScheduleInfoContextType>([
  {
    climateData: [],
    m304TimeScheduleMap: new Map<number, TimeScheduleResponse>(),
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
      climateData: [],
      m304TimeScheduleMap: new Map<number, TimeScheduleResponse>(),
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
