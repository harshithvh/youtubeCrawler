import React from "react";
import { Box, Typography } from "@mui/material";
import ErrorOutlineIcon from '@mui/icons-material/ErrorOutline';
import NoVideosImage from "../assets/error.png";

const Error = ({ message }) => (
  <Box textAlign="center" marginTop="250px">
    <ErrorOutlineIcon color="error" fontSize="large" />
    <Typography variant="h6" color="error">
      {message}
    </Typography>
    {/* <img src={NoVideosImage} alt="No Videos Found" style={{ maxWidth: "100%" }} /> */}
  </Box>
);

export default Error;
