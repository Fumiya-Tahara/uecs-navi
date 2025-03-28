import { Box, Button, Divider, Modal, Typography } from "@mui/material";
import { useState } from "react";
import { modalStyle } from "@/styles/modal";

interface ConfirmButtonProps {
  buttonLabel: string; // ボタンのラベル（例: "保存する"）
  modalTitle: string; // モーダルのタイトル（例: "ワークフローを保存しますか？"）
  confirmLabel?: string; // 確認ボタンのラベル（デフォルト: "OK"）
  cancelLabel?: string; // キャンセルボタンのラベル（デフォルト: "キャンセル"）
  onConfirm: () => void; // 確認ボタンが押されたときの処理
  buttonColor?: "primary" | "secondary" | "error"; // ボタンの色
  buttonVariant?: "contained" | "outlined"; // ボタンの塗りつぶし
}

export function ConfirmButton(props: ConfirmButtonProps) {
  const {
    buttonLabel,
    modalTitle,
    confirmLabel = "OK",
    cancelLabel = "キャンセル",
    onConfirm,
    buttonColor = "primary",
    buttonVariant = "contained",
  } = props;
  const [modalOpen, setModalOpen] = useState(false);

  const handleOpen = () => setModalOpen(true);
  const handleClose = () => setModalOpen(false);

  const handleConfirm = () => {
    onConfirm();
    setModalOpen(false);
  };

  return (
    <div>
      <Button
        variant={buttonVariant}
        size="small"
        color={buttonColor}
        onClick={handleOpen}
      >
        {buttonLabel}
      </Button>
      <Modal
        open={modalOpen}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            {modalTitle}
          </Typography>
          <Divider />
          <Box sx={{ mt: 1, display: "flex", gap: 1 }}>
            <Button variant="contained" size="small" onClick={handleConfirm}>
              {confirmLabel}
            </Button>
            <Button variant="outlined" size="small" onClick={handleClose}>
              {cancelLabel}
            </Button>
          </Box>
        </Box>
      </Modal>
    </div>
  );
}
