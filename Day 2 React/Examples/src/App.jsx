import Greetings from "./Greetings";
import Navbar from "./Navbar";
export default function APP(){
  return(
    <>
    <Navbar/>
    <h1><marquee>Welcome to REACT</marquee></h1>
    <Greetings/>
    </>
  );
}