import { createBrowserRouter, RouterProvider } from "react-router-dom";
import SettingDevice from "./routes/setting-device";
import WorkflowControl from "./routes/workflow-control";
import Root from "./routes/root";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
  },
  {
    path: "/devices",
    element: <SettingDevice />,
  },
  {
    path: "/workflow",
    element: <WorkflowControl />,
  },
]);

export const AppRouter = () => {
  return <RouterProvider router={router} />;
};