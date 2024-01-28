import axios from 'axios';

export const BASE_URL = 'http://localhost:8080';

export const fetchFromAPI = async (url) => {
  const { data } = await axios.get(`${BASE_URL}/${url}`);

  return data;
};