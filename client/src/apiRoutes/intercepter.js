import axios from "axios";
const instance = axios.create();

instance.interceptors.response.use(
  (successRes) => {
   
    if(successRes.status === 200){
        console.log(successRes.data)
        return successRes.data
    }
    
  },
  (error) => {
    if (
      error.response.status === 401 &&
      error.response.statusText === "Unauthorized"
    ) {
    //   store.dispatch(logout(true));
    }

    return Promise.reject(error);
  }
);

export default instance;