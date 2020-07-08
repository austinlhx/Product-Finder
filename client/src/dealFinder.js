import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import {
  Grid,
  Link,
} from "@material-ui/core";

import {NotificationContainer, NotificationManager} from 'react-notifications';
import 'react-notifications/lib/notifications.css';

let endpoint = "http://localhost:8080";

//let endpoint = "http://localhost:8080";

class dealFinder extends Component {
  constructor(props) {
    super(props);
    this.state = {
      open: false,
      setOpen: false,
      ProductName: "",
      ProductType: "",
      LowerBound: "",
      UpperBound: "",
      products: [],
    };
  }

  handleChange = (event) => {
    this.setState({
      [event.target.name]: event.target.value,
    });
  };

  handleSubmit = () => {
    const { ProductName, LowerBound, UpperBound, ProductType } = this.state;
    NotificationManager.success('Searching');
    
    axios
      .post(
        endpoint + "/api",
        {
          ProductName: ProductName,
          LowerBound: LowerBound,
          UpperBound: UpperBound,
          ProductType: ProductType,
        },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )
      .then((res) => {
        
        axios
          .get(endpoint + "/api", {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          })
          .then((res) => {
            if (res.data) {
              this.setState({
                products: res.data.map((item) => {
                  console.log(item);
                  return (
                    <div className="row">
                      <div className="column">
                        <img
                          src={item.Image}
                          style={{ width: "100%", height: "100%" }}
                          alt={item.Name + " image"}
                        ></img>
                      </div>
                      <div className="column">
                        <Link color="inherit" href={item.Link}>
                          {item.Name}
                        </Link>
                        <p>{"$" + item.Price}</p>
                      </div>
                    </div>
                  );
                }),
              });
            }
          });
      });
  };

  handleClose = () => {
    this.setState.setOpen = false;
  };
  handleOpen = () => {
    this.setState.setOpen = true;
  };

  render() {
    return (
      <div>
        <NotificationContainer/>
        <ul class="box-area">
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
        </ul>
        <div className="App">
          <div className="container">
            <div class="product left">
            
              <Grid container spacing={3}>
                <Grid item xs={12}>
                  <h1>Product Search</h1>
                </Grid>
                <Grid item xs={12}>
                <div class="form__group field">
                  <input
                    type="input"
                    class="form__field"
                    id="ProductName"
                    name="ProductName"
                    onChange={this.handleChange}
                    required
                  />
                  <label for="name" class="form__label">
                    Product Item
                  </label>
                </div>
                </Grid>
                <Grid item xs={12}>
                <div class="form__group field">
                  <input
                    type="input"
                    class="form__field"
                    id="ProductType"
                    name="ProductType"
                    onChange={this.handleChange}
                    required
                  />
                  <label for="name" class="form__label">
                    Product Type
                  </label>
                </div>
                </Grid>
                <Grid item xs={12}>
                <div class="form__group field">
                  <input
                    type="input"
                    class="form__field"
                    id="LowerBound"
                    name="LowerBound"
                    onChange={this.handleChange}
                    required
                  />
                  <label for="name" class="form__label">
                    Lower Bound Price
                  </label>
                </div>
                </Grid>
                <Grid item xs={12}>
                <div class="form__group field">
                  <input
                    type="input"
                    class="form__field"
                    id="UpperBound"
                    name="UpperBound"
                    onChange={this.handleChange}
                    required
                  />
                  <label for="name" class="form__label">
                    Upper Bound Price
                  </label>
                </div>
                </Grid>
                
                
                <Grid item xs={12}>
                  <button onClick={this.handleSubmit} class="searchButton">
                    Search
                  </button>
                </Grid>
              </Grid>
            </div>
            <div class="product right">
            
              {this.state.products}
              </div>
          </div>
          
        </div>
      </div>
    );
  }
}
export default dealFinder;
