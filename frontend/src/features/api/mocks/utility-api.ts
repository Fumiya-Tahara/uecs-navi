import { M304ID, ClimateData } from "@/types/api";

export function getM304s(): M304ID[] {
  const m304IDs: M304ID[] = [1, 2, 3];

  return m304IDs;
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
