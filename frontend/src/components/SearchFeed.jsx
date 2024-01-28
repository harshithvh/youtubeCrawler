import { useState, useEffect } from "react";
import { Typography, Box, IconButton } from "@mui/material";
import { useLocation } from 'react-router-dom';
import LoopIcon from '@mui/icons-material/Loop';

import { fetchFromAPI } from "../utils/fetchFromAPI";
import { Videos, Error, Pagination } from "./";

const SearchFeed = ({ live, sort, pageSize, selectedCategory }) => {
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const searchTerm = searchParams.get('query');
  const category = searchParams.get('category');

  if (category) {
    selectedCategory = category;
  }

  const [videos, setVideos] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [error, setError] = useState(null);

  const fetchData = () => {
    setVideos(null);
    setError(null);
    console.log(selectedCategory);
    fetchFromAPI(`search?query=${searchTerm}&category=${selectedCategory}&page=${currentPage}&pageSize=${pageSize}&sort=${sort}`)
      .then((data) => setVideos(data))
      .catch((error) => setError(error));
  }

  useEffect(() => {
    setVideos(null);
    fetchData();
  }, [currentPage]);


  useEffect(() => {
    setVideos(null);
    setCurrentPage(1);
    fetchData();
  }, [searchTerm, pageSize, sort, live]);

  return (
    <Box p={2} sx={{ overflowY: "auto", height: "90vh", flex: 2 }}>
      <Typography variant="h4" fontWeight={900}  color="white" mb={3} ml={{ sm: "100px"}}>
        Search Results for <span style={{ color: "#FC1503" }}>{searchTerm}</span> videos
        {live && (<IconButton color="primary" aria-label="Load" onClick={fetchData}>
                      <LoopIcon fontSize="large" />
                    </IconButton>
        )}
      </Typography>
      <Box display="flex" sx={{ flexDirection: 'column', justifyContent: !error ? 'unset' : 'center', marginLeft: "20px"}}>
        {error ? (
          <Error message={error.message}/>
          ) : (
            <>
              {<Videos videos={videos?.videos} />}
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
  </Box>
  );
};

export default SearchFeed;