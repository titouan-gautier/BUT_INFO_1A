import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { StyleSheet, Text, View, ToastAndroid, TouchableOpacity, Button, Pressable } from 'react-native';

export default class App extends React.Component {
  
  state = {
    value: 0
  }
    
  increase = () => {
    this.setState({
      value : this.state.value + 1
    })
    console.log(this.state.value)
    
  }

  decrease = () => {
    this.setState({
      value : this.state.value - 1
    })
    console.log(this.state.value)
  }

  render () {
    return (
      <View style={styles.container}>
        <Text style={styles.text}>{this.state.value}</Text>
        <StatusBar style="auto" />
        <View style={{flexDirection: 'row'}}>
          <Pressable style={styles.button} onPress={() => this.increase()} title="Increase">
            <Text style= {{color: "#fff"}}>INCREASE</Text>
          </Pressable>
          <Pressable style={styles.button} onPress={() => this.decrease()} title="Decrease" >
          <Text style= {{color: "#fff"}}>DECREASE</Text>
          </Pressable>
        </View>
      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: 'center',
    justifyContent: 'center',
  },

  text : {
    color: "#4287f5",
    fontSize: 30
  },

  button : {
    backgroundColor : "#4287f5",
    padding: 20,
    borderRadius : 15,
    marginTop: 30,
    margin: 20
  }
});
