import { BrowserRouter, Routes, Route, useLocation } from "react-router-dom";
import { Box } from '@mui/material';
import { useState, useEffect } from 'react';

import { SearchFeed, Navbar, Feed, Ping } from './components';

const App = () => {
  return (
    <BrowserRouter>
      <AppContent />
    </BrowserRouter>
  );
};

const AppContent = () => {
  const [live, setLive] = useState(false);
  const [sort, setSort] = useState("-1");
  const [pageSize, setPageSize] = useState("10");
  const [selectedCategory, setSelectedCategory] = useState("New");

  const location = useLocation();
  const isRootRoute = location.pathname === '/';

  return (
    <Box sx={{ backgroundColor: '#000' }}>
      {!isRootRoute && <Navbar live={live} setLive={setLive} sort={sort}
      setSort={setSort} pageSize={pageSize} setPageSize={setPageSize} selectedCategory={selectedCategory} />}
      <Routes>
        <Route exact path='/' element={<Ping />} />
        <Route path='/videos' element={<Feed live={live} sort={sort} pageSize={pageSize} selectedCategory={selectedCategory} 
        setSelectedCategory={setSelectedCategory}/>} />
        <Route path='/search' element={<SearchFeed live={live} sort={sort} pageSize={pageSize}  selectedCategory={selectedCategory} />} />
      </Routes>
    </Box>
  );
};

export default App;
