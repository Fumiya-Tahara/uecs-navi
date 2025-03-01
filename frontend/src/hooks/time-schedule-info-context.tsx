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
} | null;

type timeScheduleInfoContextType = [
  timeScheduleInfoState,
  Dispatch<SetStateAction<timeScheduleInfoState>>
];

const timeScheduleInfoContext = createContext<timeScheduleInfoContextType>([
  null,
  () => {},
]);

export const TimeScheduleInfoProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [timeScheduleInfo, setTimeScheduleInfo] =
    useState<timeScheduleInfoState>(null);

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
