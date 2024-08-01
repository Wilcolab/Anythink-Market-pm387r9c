import React from "react";
import logo from "../../imgs/logo.png";
import agent from "../../agent";

const Banner = (props) => {
  const onSearchChange = async (event) =>{
    
    event.preventDefault();

    props.onSearchFilter(
      event.target.value,
      (page) => 
      agent.Items.byTitle(event.target.value),
      agent.Items.byTitle(event.target.value),
    )
  }
  return (
    <div className="banner text-white">
      <div className="container p-4 text-center">
        <img src={logo}/>
        <div>
          <span>A place to </span>
          <span id="get-part">get</span>
          <form>
            <input style={{width: "260px"}} type="text" placeholder="What is that you truely desire ?" id="search-box" name="term" onChange={onSearchChange}/>          
          </form>
          <span> the cool stuff.</span>
        </div>
      </div>
    </div>
  );
};

export default Banner;
