import React, { useState, useEffect } from 'react';
import { Box, IconButton, Typography } from "@mui/material";
import { useNavigate} from 'react-router-dom';
import pongImage from '../assets/200ok.png';
import Animate from '../utils/Animate';
import { fetchFromAPI } from "../utils/fetchFromAPI";
import RocketLaunchIcon from '@mui/icons-material/RocketLaunch';

const Ping = () => {

  const [pongResponse, setPongResponse] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchFromAPI('ping')
      .then((data) => setPongResponse(data.message));
  }, []);

  const handleButtonClick = () => {
    navigate('/videos');
  };

  return (
    <>
      {pongResponse === 'PONG!' ? (
        <Box
          position="relative"
          height="100vh"
          sx={{ "::-webkit-scrollbar": { display: "none" } }}
        >
          <Animate type="fade" delay={1}>
            <Box sx={{
              position: "absolute",
              right: 0,
              height: "100%",
              width: "70%",
              backgroundPosition: "center",
              backgroundSize: "cover",
              backgroundRepeat: "no-repeat",
              backgroundImage: `url(${pongImage})`
            }} />
          </Animate>

          <Box
            position="absolute"
            left={0}
            height="100%"
            width="30%"
            bgcolor="rgba(0,0,0,0.5)"
            sx={{ display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}
          >
            <Animate type="fade" delay={0.5}>
              <Typography variant="h2" color="white" fontWeight="bold" mb={2}>
                {pongResponse}
              </Typography>
            </Animate>

            <Animate type="fade" delay={1}>
              <IconButton color="primary" aria-label="Load" onClick={handleButtonClick}>
                <RocketLaunchIcon sx={{ fontSize: 50 }} />
              </IconButton>
            </Animate>
          </Box>
        </Box>
        ) : (
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              height: '100vh',
              backgroundColor: '#000',
              color: '#fff',
          }}
        >
          <Typography variant="h4">Waiting for server response...</Typography>
        </Box>
    )}
    </>
  );
};

export default Ping;