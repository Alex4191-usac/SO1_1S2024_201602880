import { Link } from "react-router-dom";
const Navbar = () => {


    return (
        <nav className="flex bg-slate-100">
            <div className="w-1/4 p-5 ">
                <p className="text-2xl font-bold text-slate-500">SO1-201602880</p>
            </div>
           <div className="w-3/4 flex  justify-between p-5 mr-2 cursor-pointer ">
                <Link className="text-xl" to="/">System Monitor</Link>
                <Link className="text-xl" to="history">History Ram & CPU</Link>
                <Link className="text-xl" to="tree">Tree Process</Link>
                <Link className="text-xl" to="smanagment">State Managment</Link>
           </div>
        </nav>
    )
};

export default Navbar;