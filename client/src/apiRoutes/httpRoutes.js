import axios  from "./intercepter";
const baseUrl = "http://localhost:8080";
export const getPlayer = () => {
  return axios.get(`${baseUrl}/players`);
};