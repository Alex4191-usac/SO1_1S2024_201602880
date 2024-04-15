import { Route, createBrowserRouter, createRoutesFromElements} from "react-router-dom";
import App from "../App";
import Navbar from "../components/Navbar";
import Root from "../pages/Root";
import History from "../pages/History";
import Tree from "../pages/Tree";
import SManagment from "../pages/SManagment";

export const router = createBrowserRouter(
    createRoutesFromElements (
        <>
            
            <Route path="/" element={<Root />} >
                <Route index  element={<App />} />
                <Route path="history" element={<History />} />
                <Route path="tree" element={<Tree/>}/>
                <Route path="smanagment" element={<SManagment/>}/>
            </Route>    
            <Route path="*" element={<div>NOT FOUND</div>} />
        </>
    )
);