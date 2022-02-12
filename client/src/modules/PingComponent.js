import React, { useState, useEffect } from "react";
import axios from "axios";

function PingComponent() {
  const [pingState, setPingState] = useState("");

  useEffect(() => {
    axios
      .get("api/ping")
      .then((response) => {
        setPingState(response.data.message);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, []);

  return <h1>Ping {pingState}</h1>;
}

export default PingComponent;
