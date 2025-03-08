import {
  ClimateData,
  DeviceResponse,
  WorkflowWithUIResponse,
} from "@/types/api";

export function getDevices(m304ID: number): DeviceResponse[] | undefined {
  const map = new Map<number, DeviceResponse[]>();

  const devices1: DeviceResponse[] = [
    { id: 1, climate_data_id: 1, m304_id: 1, name: "ヒーター", rly: 1 },
    { id: 2, climate_data_id: 2, m304_id: 1, name: "ミスト", rly: 2 },
    {
      id: 3,
      climate_data_id: 3,
      m304_id: 1,
      name: "二酸化炭素供給装置",
      rly: 3,
    },
    { id: 4, climate_data_id: 4, m304_id: 1, name: "照明", rly: 4 },
  ];
  const devices2: DeviceResponse[] = [
    { id: 1, climate_data_id: 1, m304_id: 2, name: "ヒーター", rly: 1 },
    { id: 2, climate_data_id: 2, m304_id: 2, name: "ミスト", rly: 2 },
    {
      id: 3,
      climate_data_id: 3,
      m304_id: 2,
      name: "二酸化炭素供給装置",
      rly: 3,
    },
    { id: 4, climate_data_id: 4, m304_id: 2, name: "照明", rly: 4 },
  ];
  const devices3: DeviceResponse[] = [
    { id: 1, climate_data_id: 1, m304_id: 3, name: "ヒーター", rly: 1 },
    { id: 2, climate_data_id: 2, m304_id: 3, name: "ミスト", rly: 2 },
    {
      id: 3,
      climate_data_id: 3,
      m304_id: 3,
      name: "二酸化炭素供給装置",
      rly: 3,
    },
    { id: 4, climate_data_id: 4, m304_id: 3, name: "照明", rly: 4 },
  ];

  map.set(1, devices1);
  map.set(2, devices2);
  map.set(3, devices3);

  return map.get(m304ID);
}

export function getClimateDatas(): ClimateData[] {
  const climateDatas: ClimateData[] = [
    { id: 1, name: "気温", unit: "℃" },
    { id: 2, name: "湿度", unit: "%" },
    { id: 3, name: "二酸化炭素量", unit: "ppm" },
    { id: 4, name: "照度", unit: "lx" },
  ];

  return climateDatas;
}

export function getWorkflowsWithUI(
  m304ID: number
): WorkflowWithUIResponse[] | undefined {
  const map = new Map<number, WorkflowWithUIResponse[]>();
  const workflowWithUIRes1: WorkflowWithUIResponse[] = [
    {
      workflow: {
        id: 1,
        m304_id: 1,
        name: "Mock Workflow1",
      },
      workflow_ui: {
        nodes: [
          {
            id: 1,
            workflow_id: 1,
            workflow_node_id: "workflow_name_1",
            node_type: "workflow_name",
            data: { name: "Mock Workflow1" },
            position_x: 100,
            position_y: 200,
          },
          {
            id: 2,
            workflow_id: 1,
            workflow_node_id: "operation_1",
            node_type: "operation",
            data: {
              device_id: 1,
            },
            position_x: 300,
            position_y: 400,
          },
          {
            id: 3,
            workflow_id: 1,
            workflow_node_id: "operation_2",
            node_type: "operation",
            data: {
              device_id: 2,
            },
            position_x: 500,
            position_y: 600,
          },
        ],
        edges: [
          {
            id: 1,
            workflow_id: 1,
            source_node_id: "workflow_name_1",
            target_node_id: "operation_1",
          },
          {
            id: 2,
            workflow_id: 1,
            source_node_id: "operation_1",
            target_node_id: "operation_2",
          },
        ],
      },
    },
    {
      workflow: {
        id: 2,
        m304_id: 1,
        name: "Mock Workflow2",
      },
      workflow_ui: {
        nodes: [
          {
            id: 4,
            workflow_id: 2,
            workflow_node_id: "workflow_name_2",
            node_type: "workflow_name",
            data: { name: "Mock Workflow2" },
            position_x: 100,
            position_y: 200,
          },
        ],
        edges: [],
      },
    },
  ];

  const workflowWithUIRes2: WorkflowWithUIResponse[] = [
    {
      workflow: {
        id: 3,
        m304_id: 2,
        name: "Mock Workflow3",
      },
      workflow_ui: {
        nodes: [
          {
            id: 5,
            workflow_id: 3,
            workflow_node_id: "workflow_name_3",
            node_type: "workflow_name",
            data: { name: "Mock Workflow3" },
            position_x: 100,
            position_y: 200,
          },
        ],
        edges: [],
      },
    },
  ];

  const workflowWithUIRes3: WorkflowWithUIResponse[] = [
    {
      workflow: {
        id: 4,
        m304_id: 3,
        name: "Mock Workflow4",
      },
      workflow_ui: {
        nodes: [
          {
            id: 6,
            workflow_id: 4,
            workflow_node_id: "workflow_name_4",
            node_type: "workflow_name",
            data: { name: "Mock Workflow5" },
            position_x: 100,
            position_y: 200,
          },
        ],
        edges: [],
      },
    },
  ];

  map.set(1, workflowWithUIRes1);
  map.set(2, workflowWithUIRes2);
  map.set(3, workflowWithUIRes3);

  return map.get(m304ID);
}

export function deleteWorkflowWithUI(workflowID: number) {
  console.log("Deleted: ", workflowID);
}
