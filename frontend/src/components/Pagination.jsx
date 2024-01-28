import React from "react";
import { Pagination as MuiPagination } from "@mui/material";

const Pagination = ({ currentPage, pageSize, totalCount, setCurrentPage, totalPages }) => {

  const handleChange = (event, value) => {
    setCurrentPage(value)
  };


  return (
    <MuiPagination
      count={totalPages}
      page={currentPage}
      onChange={handleChange}
      color="primary"
      size="large"
      sx={{ mt: 3, mb: 3, display: "flex", justifyContent: "center", 
        "& .MuiPaginationItem-icon": {
            color: "#fff",
        },"& .MuiPaginationItem-page": {
            color: "#fff",
        }, }}
    />
  );
};

export default Pagination;
