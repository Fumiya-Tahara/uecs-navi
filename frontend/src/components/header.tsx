import {
  AppBar,
  Toolbar,
  Box,
  Link,
  List,
  ListItem,
  ListItemButton,
} from "@mui/material";
import HomeIcon from "@mui/icons-material/Home";
import SettingsRemoteIcon from "@mui/icons-material/SettingsRemote";
import TimelineIcon from "@mui/icons-material/Timeline";
import AccessTimeIcon from "@mui/icons-material/AccessTime";
import Image from "next/image";

interface HeaderItem {
  title: string;
  path: string;
  icon: JSX.Element;
}

const headerItems: HeaderItem[] = [
  {
    title: "ハウス状態",
    path: "/",
    icon: <HomeIcon />,
  },
  {
    title: "デバイス設定",
    path: "/devices",
    icon: <SettingsRemoteIcon />,
  },
  {
    title: "ワークフロー制御",
    path: "/workflows",
    icon: <TimelineIcon />,
  },
  {
    title: "スケジュール設定",
    path: "/schedules",
    icon: <AccessTimeIcon />,
  },
];

export default function Header() {
  return (
    <AppBar
      position="fixed"
      sx={{
        zIndex: (theme) => theme.zIndex.drawer + 1,
        backgroundColor: "primary.main",
      }}
    >
      <Toolbar
        sx={{ height: "70px", backgroundColor: "#DDEAF6", color: "#000" }}
      >
        <Box
          sx={{
            width: "100%",
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
          }}
        >
          <Box
            sx={{
              paddingTop: "8px",
              width: "120px",
            }}
          >
            <Link href="/" underline="none">
              <Image
                src="/uecs_navi_logo.png"
                alt="UECS Navi Logo"
                width={160}
                height={75}
                layout="responsive"
                objectFit="contain"
              />
            </Link>
          </Box>
          <Box>
            <List sx={{ display: "flex", fontSize: "12px" }}>
              {headerItems.map((item, index) => (
                <ListItem key={index} sx={{ width: "auto" }} disablePadding>
                  <ListItemButton sx={{ display: "flex", gap: 1 }}>
                    {item.icon}
                    <Link href={item.path} underline="none" color="inherit">
                      {item.title}
                    </Link>
                  </ListItemButton>
                </ListItem>
              ))}
            </List>
          </Box>
        </Box>
      </Toolbar>
    </AppBar>
  );
}
