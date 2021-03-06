import React from 'react';

export default class CreateTaskForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {description: 'Default task',
                    status: 'todo'};
    
        this.handleChangeText = this.handleChangeText.bind(this);
        this.handleChangeDropdown = this.handleChangeDropdown.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }
    
      handleChangeText(event) {
        this.setState({description: event.target.value});
      }

      handleChangeDropdown(event) {
        this.setState({status: event.target.value});
      }
    
      handleSubmit() {
	console.log(this.state);
        //alert('A task added: ' + this.state.status);
        fetch("http://test.khanhquocnguyen.com:8080/tasks", {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(this.state)})
        //.then(response => response.json())
        .then(data => {
          console.log('Success:', data);
        })
        .catch((error) => {
          console.error('Submit Error:', error);
        
        });
      }
    
      render() {
        return (
          <form onSubmit={this.handleSubmit}>
            <label>
              Add task:
              <input type="text" value={this.state.description} onChange={this.handleChangeText} />
              <select value={this.state.status} onChange={this.handleChangeDropdown}>
                <option value="todo">todo</option>
                <option value="doing">doing</option>
                </select>
            </label>
            <input type="submit" value="Submit" />
          </form>
        );
      }
}
