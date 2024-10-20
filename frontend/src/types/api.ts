export interface HouseResponse {
  id: number;
  name: string;
}

export interface JoinedDeviceResponse {
  id: number;
  name: string;
  houseID: number;
  setPoint?: number;
  duration?: number;
  climateData: string;
  unit: string;
}

export interface ClimateDataResponse {
  id: number;
  climateData: string;
  unit: string;
}

export interface CreateDeviceRequest {
  name: string;
  climateDataID: number;
  setPoint?: number;
  duration?: number;
}

export interface CreateHouseRequest {
  name: string;
}
