import React from 'react';
import './App.css';
import TaskList from './comps/TaskList';
import CreateTaskForm from './comps/CreateTaskForm'

function App() {
  return (
    <div className="App">
      <TaskList />
      <CreateTaskForm />
    </div>
  );
}

export default App;
