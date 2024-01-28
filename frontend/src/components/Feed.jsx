import React, { useEffect, useState } from "react";
import { Box, Stack, Typography, IconButton } from "@mui/material";
import LoopIcon from '@mui/icons-material/Loop';

import { fetchFromAPI } from "../utils/fetchFromAPI";
import { Videos, Sidebar, Pagination, Error } from "./";

const Feed = ({ live, sort, pageSize, selectedCategory, setSelectedCategory }) => {

  const [videos, setVideos] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [error, setError] = useState(null);

  const fetchData = () => {
    setVideos(null);
    setError(null);
    fetchFromAPI(`videos?category=${selectedCategory}&page=${currentPage}&pageSize=${pageSize}&sort=${sort}`)
      .then((data) => setVideos(data))
      .catch((error) => setError(error));
  };

  useEffect(() => {
    setVideos(null);
    fetchData();
  }, [currentPage]);

  useEffect(() => {
    setVideos(null);
    setCurrentPage(1);
    fetchData();
  }, [selectedCategory, live, pageSize, sort]);

  return (
    <Stack sx={{ flexDirection: { sx: "column", md: "row" } }}>
      <Box sx={{ height: { sx: "auto", md: "92vh" }, borderRight: "1px solid #3d3d3d", px: { sx: 0, md: 2 } }}>
        <Sidebar selectedCategory={selectedCategory} setSelectedCategory={setSelectedCategory} />
      </Box>

      <Box p={2} sx={{ overflowY: "auto", height: "90vh", flex: 2 }}>
        <Typography variant="h4" fontWeight="bold" mb={2} sx={{ color: "white" }}>
          {selectedCategory} <span style={{ color: "#FC1503" }}>videos</span>
          {live && (<IconButton color="primary" aria-label="Load" onClick={fetchData}>
                      <LoopIcon fontSize="large" />
                    </IconButton>
          )}
        </Typography>

        {error ? (
          <Error message={error.message}/>
          ) : (
          <>
            <Videos videos={videos?.videos} />

            {videos && (
              <Pagination
                currentPage={currentPage}
                pageSize={pageSize}
                totalCount={videos.pagination.totalVideos}
                setCurrentPage={setCurrentPage}
                totalPages={videos.pagination.totalPages}
              />
            )}
          </>
        )}
      </Box>
    </Stack>
  );
};

export default Feed;