import React, { Component } from "react";
import axios from "axios";
import "./App.css";
import {
  FormControl,
  Input,
  InputLabel,
  FormHelperText,
  MenuItem,
  Select,
  Typography,
  Grid,
  Button,
  isMuiElement,
} from "@material-ui/core";

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
    console.log(this.state);
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
                          style={{ width: "100%" }}
                        ></img>
                      </div>
                      <div className="column">
                        <a href={item.Link}>{item.Name}</a>
                        <p>{item.Price}</p>
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
                  <Typography variant="h4" color="textSecondary">
                    Product Search
                  </Typography>
                </Grid>
                <Grid item xs={12}>
                  <FormControl>
                    <InputLabel htmlFor="my-input">Product Item</InputLabel>
                    <Input
                      id="ProductName"
                      name="ProductName"
                      aria-describedby="my-helper-text"
                      onChange={this.handleChange}
                    />
                    <FormHelperText id="my-helper-text">
                      e.g. Airpods Pro, Beats Headphones, etc.
                    </FormHelperText>
                  </FormControl>
                </Grid>
                <Grid item xs={12}>
                  <FormControl>
                    <InputLabel id="demo-controlled-open-select-label">
                      Product Type
                    </InputLabel>
                    <Select
                      labelId="demo-controlled-open-select-label"
                      open={this.open}
                      onClose={this.handleClose}
                      onOpen={this.handleOpen}
                      onChange={this.handleChange}
                      name="ProductType"
                    >
                      <MenuItem value={"electronics"}>Electronics</MenuItem>
                      <MenuItem value={"clothing"}>Clothing</MenuItem>
                      <MenuItem value={"apparel"}>Apparel</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12}>
                  <FormControl>
                    <InputLabel htmlFor="my-input">
                      Lower Bound Price
                    </InputLabel>
                    <Input
                      aria-describedby="my-helper-text"
                      id="LowerBound"
                      name="LowerBound"
                      onChange={this.handleChange}
                    />
                    <FormHelperText>e.g. 15</FormHelperText>
                  </FormControl>
                </Grid>
                <Grid item xs={12}>
                  <FormControl>
                    <InputLabel htmlFor="my-input">
                      Upper Bound Price
                    </InputLabel>
                    <Input
                      id="UpperBound"
                      name="UpperBound"
                      onChange={this.handleChange}
                      aria-describedby="my-helper-text"
                    />
                    <FormHelperText id="my-helper-text">e.g. 18</FormHelperText>
                  </FormControl>
                </Grid>
                <Grid item xs={12}>
                  <Button onClick={this.handleSubmit}>Search</Button>
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
