import { ClimateData, Workflow } from "@/types/api";
import {
  createContext,
  useContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

type timeScheduleInfoState = {
  climate_data: ClimateData[];
  m304WorkflowMap: Map<number, Workflow[]>;
};

type timeScheduleInfoContextType = [
  timeScheduleInfoState,
  Dispatch<SetStateAction<timeScheduleInfoState>>
];

const timeScheduleInfoContext = createContext<timeScheduleInfoContextType>([
  {
    climate_data: [],
    m304WorkflowMap: new Map<number, Workflow[]>(),
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
      m304WorkflowMap: new Map<number, Workflow[]>(),
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
