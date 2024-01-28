import { Stack, IconButton, Button } from '@mui/material';
import { Link } from 'react-router-dom';
import YTLogo from '../assets/fam.jpeg';
import { SearchBar } from './';
import { fetchFromAPI } from "../utils/fetchFromAPI";
import { ArrowUpward, ArrowDownward, Stop } from '@mui/icons-material';
import CircleIcon from '@mui/icons-material/Circle';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';

const Navbar = ({ live, setLive, sort, setSort, pageSize, setPageSize, selectedCategory }) => {

    const InitBackgroundFetch = () => {
        fetchFromAPI(`start-fetching?category=${selectedCategory}`)
          .catch((error) => setError(error));
    };

    const StopBackgroundFetch = () => {
        fetchFromAPI(`stop-fetching`)
          .catch((error) => setError(error));
    };

    const handleToggleAscending = () => {
        if (sort === '1') {
            setSort('-1');
        } else {
            setSort('1');
        }
    };

	const handleLiveToggle = () => {
        setLive(!live);
        if (!live) {
            InitBackgroundFetch();
        } else {
            StopBackgroundFetch();
        }
    };

    return (
        <Stack
            direction='row'
            alignItems='center'
            p={2}
            sx={{ position: 'sticky', background: '#000', top: 0, justifyContent: 'center', gap: '10px', marginLeft: '290px' }}
        >
            <Link to='/videos' style={{ display: 'flex', alignItems: 'center', borderRadius: '10px', marginRight: '15px' }}>
                <img src={YTLogo} alt='logo' height={45} style={{ borderRadius: '10px' }} />
            </Link>
            <SearchBar selectedCategory={selectedCategory}/>
            <Stack direction="row" alignItems="center" spacing={5} sx={{ marginLeft: '15px' }}>
                {live ? (
                    <Button variant="contained" endIcon={<Stop />} onClick={handleLiveToggle}>
                        Stop
                    </Button>
                ) : (
                    <Button variant="contained" color="error"  onClick={handleLiveToggle}>
                        <span style={{ marginRight: '9px' }}>Live</span>
						<CircleIcon  fontSize='5px'/>
                    </Button>
                )}
				<Select
					variant="plain"
					value={pageSize}
					onChange={(_, value) => setPageSize(value)}
					slotProps={{
						listbox: {
						variant: 'outlined',
						},
					}}
					sx={{ mr: -1.5, '&:hover': { bgcolor: 'darkgrey' } }}
					>
					<Option value="1">1</Option>
					<Option value="2">2</Option>
					<Option value="3">3</Option>
					<Option value="4">4</Option>
					<Option value="5">5</Option>
					<Option value="6">6</Option>
					<Option value="7">7</Option>
					<Option value="8">8</Option>
					<Option value="9">9</Option>
					<Option value="10">10</Option>
					<Option value="11">11</Option>
					<Option value="12">12</Option>
					<Option value="13">13</Option>
					<Option value="14">14</Option>
					<Option value="15">15</Option>
				</Select>
                <IconButton onClick={handleToggleAscending} sx={{ color: '#fff' }}>
                    {sort === "1" ? <ArrowUpward /> : <ArrowDownward />}
                </IconButton>
            </Stack>
        </Stack>
    );
};

export default Navbar;

