// This file was added by Matt
// It contains the component for a New Announcement modal item 
// which is dispayed when a user wants to make a new announcement
// it works like a popup.

import React from 'react';
import {Alert, Modal, StyleSheet, Pressable, View, Dimensions, TextInput} from 'react-native';
import { Block, Text, theme } from 'galio-framework';
import { Images, argonTheme } from "../constants";

const { width, height } = Dimensions.get("screen");

class NewAnouncementModal extends React.Component {
  state = {
    modalVisible: false,
  };

  render() {
    const {modalVisible} = this.state;
    return (
      <View style={styles.centeredView}>
        
        <Modal
          animationType="fade"
          transparent={true}
          visible={modalVisible}
          onRequestClose={() => {
            Alert.alert('Modal has been closed.');
            this.setState({modalVisible: !modalVisible});
          }}>
          <View style={styles.centeredView}>    
            <View style={styles.modalView}>
              <Block style={styles.header}>

                <Text style={styles.header}>
                  New Announcement 
                  <Pressable
                      onPress={() => this.setState({modalVisible: !modalVisible})}>
                      <Text style={styles.exit}>x</Text>
                  </Pressable>
                </Text>
              </Block>

              <Text style={styles.secondaryHeader}>
                Post title
              </Text>

              <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                <TextInput
                  style={styles.input} 
                  placeholder="Enter a title"
                  placeholderTextColor={"lightgrey"}                  
                  // onChangeText={handleInput}
                />
              </Block>

              <Text style={styles.secondaryHeader}>
                Description
              </Text>
              <Block width={width * 0.8 - 20}  style={{ marginBottom: 15 }}>
                <TextInput
                  style={styles.input}
                  placeholder="Provide announcement details here"
                  placeholderTextColor={"lightgrey"}
                  height={height * 0.3}
                  textAlignVertical={'top'}
                  paddingTop={10}
                  multiline={true}
                  // onChangeText={handleInput}
                />
              </Block>


              <Pressable
                style={[styles.button, styles.buttonClose]}
                onPress={() => this.setState({modalVisible: !modalVisible})}>
                <Text style={styles.textStyle}>CREATE ANNOUNCEMENT</Text>
              </Pressable>

            </View>
          </View>
        </Modal>
        {/* button that shows before opening modal */}
        <Pressable
          style={[styles.button, styles.buttonOpen]}
          onPress={() => this.setState({modalVisible: true})}>
          <Text style={styles.textStyle}>New Announcement</Text>
        </Pressable>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  exit: {
    // borderRadius: 20,
    // padding: 10,
    fontWeight: 'bold',
    elevation: 2,
    
  },
  header: {
    fontSize: 25,
    fontWeight: 'bold',
    color: "#525F7F",
    marginBottom: 10,
    marginTop: 10,
  },
  secondaryHeader: {
    fontSize: 17,
    fontWeight: 'bold',
    color: "#525F7F",
    marginBottom:5,
  },
  
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    borderWidth: .5,
    borderRadius: 5,
    // borderColor: "lightgrey",
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
   
  },

  // matt's added styles above ^^^


  centeredView: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 22,
  },
  modalView: {
    margin: 20,
    backgroundColor: 'white',
    // borderRadius: 5,
    padding: 35,
    paddingTop: 0,
    // alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 4,
    elevation: 5,
  },
  button: {
    borderRadius: 5,
    padding: 10,
    elevation: 2,
  },
  buttonOpen: {
    backgroundColor: '#F194FF',
  },
  buttonClose: {
    backgroundColor: '#5e72e4',
    // matt---
    // borderRadius: 5,
    padding: 10,
    marginTop:10,
    // matt---
  },
  textStyle: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  modalText: {
    marginBottom: 15,
    textAlign: 'center',
  },
});


export default NewAnouncementModal;