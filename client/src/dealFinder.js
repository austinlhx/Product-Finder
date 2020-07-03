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
} from "@material-ui/core";

//let endpoint = "http://localhost:8080";

class dealFinder extends Component {
  constructor(props) {
    super(props);
    this.state = {
      age: "",
      setAge: "",
      open: false,
      setOpen: false,
      ProductName: "",
      LowerBound: "",
      UpperBound: "",
    };
  }

  handleChange = (event) => {
    this.setState({
        [event.target.name]: event.target.value,
      });
  };

  Submit = () => {
      const {
          ProductName: ProductName,
      } = this.state;
      console.log(ProductName)
      
      console.log("button clicked")
      console.log(this.state)
  }

  
  handleClose = () => {
    this.setState.setOpen = false;
  };
  handleOpen = () => {
    this.setState.setOpen = true;
  };
  render() {
      const {
          ProductName,
      } = this.state;
    return (
      <div className="App">
        <div className="container">
          <div class="product left">
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Typography variant="h4">Product Search</Typography>
              </Grid>
              <Grid item xs={12}>
                <FormControl>
                  <InputLabel htmlFor="my-input" onChange={this.handleChange}
                  value={ProductName || ""}>Product Item</InputLabel>
                  <Input id="my-input" aria-describedby="my-helper-text" />
                  <FormHelperText id="my-helper-text">
                    e.g. Airpods Pro, Beats Headphones, etc.
                  </FormHelperText>
                </FormControl>
              </Grid>
              <Grid item xs={12}>
                <FormControl>
                  <InputLabel htmlFor="my-input">Product Item</InputLabel>
                  <Input id="my-input" aria-describedby="my-helper-text" />
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
                    id="demo-controlled-open-select"
                    open={this.open}
                    onClose={this.handleClose}
                    onOpen={this.handleOpen}
                    value={this.age}
                    onChange={this.handleChange}
                  >
                    <MenuItem value={"electronics"}>Electronics</MenuItem>
                    <MenuItem value={"clothing"}>Clothing</MenuItem>
                    <MenuItem value={"apparel"}>Apparel</MenuItem>
                  </Select>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl>
                  <InputLabel htmlFor="my-input">Lower Bound Price</InputLabel>
                  <Input id="my-input" aria-describedby="my-helper-text" />
                  <FormHelperText id="my-helper-text">e.g. 15</FormHelperText>
                </FormControl>
              </Grid>
              <Grid item xs={6}>
                <FormControl>
                  <InputLabel htmlFor="my-input">Upper Bound Price</InputLabel>
                  <Input id="my-input" aria-describedby="my-helper-text" />
                  <FormHelperText id="my-helper-text">e.g. 18</FormHelperText>
                </FormControl>
              </Grid>
              <Grid item xs={12}>
                <Button onClick = {this.Submit}>Search</Button>
              </Grid>
            </Grid>
          </div>
          <div class="product right">
            <Typography variant="h4">Product Search</Typography>
            <p>some</p>
          </div>
        </div>
      </div>
    );
  }
}
export default dealFinder;
