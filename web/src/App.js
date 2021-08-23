import "./App.css";
import TextField from "@material-ui/core/TextField";
import { React, Component } from "react";
import Button from "@material-ui/core/Button";
import { createTheme } from "@material-ui/core/styles";
import { ThemeProvider } from "@material-ui/styles";

const theme = createTheme({
  palette: {
    secondary: {
      // This is green.A700 as hex.
      main: "#11cb5f",
    },
  },
});

class App extends Component {
  state = {
    dishName: "",
    noOfIngredients: 0,
    ingredientsDetails: "",
    preparationProcess: "",
    buttoncolortgt: "primary",
  };

  handleDishName = (event) => {
    this.setState({ dishName: event.target.value });
  };

  handleNoOfIngredients = (event) => {
    this.setState({ noOfIngredients: event.target.value });
  };

  handleIngredientsDetails = (event) => {
    this.setState({ ingredientsDetails: event.target.value });
  };

  handlePreparationProcess = (event) => {
    this.setState({ preparationProcess: event.target.value });
  };

  onCreateButtonSubmit = () => {
    const name = this.state.dishName;
    const dishname = name.charAt(0).toUpperCase() + name.slice(1);

    const ingDetailsRaw = this.state.ingredientsDetails;

    const arr = ingDetailsRaw.split(" ");
    const mapIngDetails = {};

    for (let i = 0; i < arr.length - 1; i += 2) {
      mapIngDetails[arr[i]] = arr[i + 1];
    }

    const data = {
      name: dishname,
      noofingredients: parseInt(this.state.noOfIngredients),
      preparation: this.state.preparationProcess,
      ingredientsdetails: mapIngDetails,
    };

    fetch("http://localhost:8070/recipe/create/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((resp) => resp.json())
      .then((resp) => {
        console.log(resp, resp.name, data.name);
        if (resp.name === data.name) {
          this.setState({ buttoncolortgt: "secondary" });
        }
      });
  };

  render() {
    return (
      <div className="App">
        <div className="createrecipe">Create a recipe</div>
        <div className="wrapper">
          <div className="childwrapper">
            <ThemeProvider theme={theme}>
              <Button
                style={buttonStyleCreate}
                variant="contained"
                color={this.state.buttoncolortgt}
                onClick={this.onCreateButtonSubmit}
              >
                Create
              </Button>
            </ThemeProvider>
            <div className="dishNoDetailsWrapper">
              <TextField
                onChange={this.handleDishName}
                style={paddingText}
                label="Dish Name"
                variant="outlined"
              />
              <TextField
                onChange={this.handleNoOfIngredients}
                style={paddingText}
                label="No. of Ingredients"
                variant="outlined"
              />
              <TextField
                onChange={this.handleIngredientsDetails}
                style={paddingText}
                rows={10}
                multiline
                label="Ingredients Details"
                variant="outlined"
              />
            </div>
          </div>
          <div className="childwrapper">
            <TextField
              onChange={this.handlePreparationProcess}
              style={paddingText}
              rows={20}
              multiline
              label="Preparation Process"
              variant="outlined"
            />
          </div>
        </div>
      </div>
    );
  }
}

export default App;

const paddingText = {
  paddingBottom: "10px",
};

const buttonStyleCreate = {
  marginBottom: "22px",
};
