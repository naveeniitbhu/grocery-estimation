import './App.css';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { React, Component } from 'react';

class App extends Component {

  state = {
    dishName: '',
    noOfIngredients: 0,
    preparationDetails:'',
    ingredientDetails: {},
  }


  onButtonCreateSubmit = () => {
    console.log("clicked create")
    console.log(this.state)
    fetch('http://localhost:8070/recipe/create/', {
      method: 'POST',
      headers: {'Content-Type' : 'application/json'},
      body: JSON.stringify({
        name: this.state.dishName,
        noofingredients: this.state.noOfIngredients,
        preparation: this.state.preparationDetails,
        ingredientsdetails: {"a":"10", "b":"20", "c":"30"}
      })
    })
  }

  handleDishName = (event) => {
    this.setState({ dishName: event.target.value })
  }

  handleNoOfIngredients = (event) => {
    this.setState({ noOfIngredients: parseInt(event.target.value) })
  }

  handlePreparationDetails = (event) => {
    this.setState({ preparationDetails: event.target.value })
  }

  handlePreparationDetails = (event) => {
    this.setState({ preparationDetails: event.target.value })
  }

  handleIngredientsDetails = (event) => {
    console.log(event.target.value)
  }
 

  render() {
    return (
      <div className="App">
        <div className="createrecipe">
          <div className="createrecipetitle">
            <span style={paddingText}>Create a Recipe</span>
            <Button variant="contained" color="primary" onClick={this.onButtonCreateSubmit}>
              Create
            </Button>
          </div>
          <div>
            <div>
              <TextField style={paddingText} label="Dish Name" variant="outlined" 
                onChange={this.handleDishName} 
              />
              <TextField style={paddingText} label="No. Of Ingredients" variant="outlined"
                onChange={this.handleNoOfIngredients}
              />
              <TextField style={paddingText} label="Ingredients Details" variant="outlined" rows={10} multiline
                onChange={this.handleIngredientsDetails}
              />
            </div>
            <div>
              <TextField style={paddingText} label="Preparation Details" variant="outlined" rows={12} multiline
                onChange={this.handlePreparationDetails}
              />
            </div>
          </div>
  
  

        </div>
        
      </div>
    );
  }
 
}

export default App;

const paddingText = {
  padding:'10px'
}
