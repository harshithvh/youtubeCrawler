import React from 'react';
import { Typography, Card, CardContent, CardMedia } from '@mui/material';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import {
	demoThumbnailUrl,
	demoVideoTitle,
	demoChannelTitle,
} from '../../utils/constants';
import './VideoCard.scss';

const calculateTimeDifference = (uploadTime) => {
	const currentTime = new Date();
	const uploadedAt = new Date(uploadTime);
	const difference = currentTime - uploadedAt;
	const seconds = Math.floor(difference / 1000);
	const minutes = Math.floor(seconds / 60);
	const hours = Math.floor(minutes / 60);
	const days = Math.floor(hours / 24);

	if (days > 0) {
	  return `${days} day${days > 1 ? 's' : ''} ago`;
	} else if (hours > 0) {
	  return `${hours} hour${hours > 1 ? 's' : ''} ago`;
	} else if (minutes > 0) {
	  return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
	} else {
	  return `${seconds} second${seconds > 1 ? 's' : ''} ago`;
	}
};

const VideoCard = ({
	video
}) => (
	<Card
		sx={{
			width: { xs: '100%', sm: '358px', md: '358px' },
			boxShadow: 'none',
			borderRadius: '10px',
			overflow: 'hidden',
		}}
		className='videoCardContainer'
	>
		<CardMedia
			image={video?.ThumbnailURL || demoThumbnailUrl}
			alt={video?.Title}
			sx={{ width: { xs: '100%', sm: '358px' }, height: 180 }}
		/>
		<CardContent sx={{ backgroundColor: '#1E1E1E', height: '106px' }}>
			<Typography variant='subtitle1' fontWeight='bold' color='#FFF'>
				{video?.Title.slice(0, 60) || demoVideoTitle.slice(0, 60)}
			</Typography>
			<Typography variant='subtitle2' color='gray'>
				{video?.ChannelTitle || demoChannelTitle}
			<CheckCircleIcon sx={{ fontSize: '12px', color: 'gray', ml: '5px' }} />
			</Typography>
			<Typography variant='subtitle2' color='gray'>
				{video?.Views || '1.2M'} views • {calculateTimeDifference(video?.PublishedAt) || '1 day ago'}
			</Typography>
		</CardContent>
	</Card>
);

export default VideoCard;
