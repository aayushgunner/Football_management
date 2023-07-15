import axios  from "./intercepter";
const baseUrl = "http://localhost:8080";
export const getPlayer = () => {
  return axios.get(`${baseUrl}/players`);
};
export const getTeam = () => {
  return axios.get(`${baseUrl}/teams`)
}
export const getStat = (player_id) => {
  return axios.get(`${baseUrl}/stats/${player_id}`)
}