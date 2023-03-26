// This file was added by Matt
// It contains the component for a New post modal item 
// which is dispayed when a user wants to make a new announcement
// it works like a popup. 

import React from 'react';
import {Alert, Modal, StyleSheet, Pressable, View, Dimensions, TextInput, Image, ScrollView} from 'react-native';
import { Block, Text, theme } from 'galio-framework';
import { argonTheme } from "../constants";
import ImagePicker from '../components//ImagePicker.js';


const { width, height } = Dimensions.get("screen");

function handleCreatePostBtnClick() {
// write function for create announcement button here


}

function uploadImageButton() {
  // write function for create announcement button here
  
  
}

class NewPostModal extends React.Component {
  state = {
    modalVisible: false,
  };


  render() {
    const {modalVisible} = this.state;
    return (

      <View style={styles.centeredView}>
        
        <Modal
          propagateSwipe={true}
          animationType="fade"
          transparent={true}
          visible={modalVisible}
          onRequestClose={() => {
            Alert.alert('Modal has been closed.');
            this.setState({modalVisible: !modalVisible});
          }}>
             
             <ScrollView> 
          <View style={styles.centeredView}>    
            <View style={styles.modalView}>
              {/* x icont to exit modal */}
              <Pressable
                onPress={() => this.setState({modalVisible: !modalVisible})}
                style={{alignItems: 'flex-end'}}>
                <Image
                style={styles.exit}
                source={require('../assets/imgs/exit.png')} />
              </Pressable>
              
              <View style={styles.modalViewInside}>
   
                <Text style={styles.header}>
                  Create New Post 
                </Text>
     

                <Text style={styles.secondaryHeader}>
                  Post title
                </Text>
                <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input} 
                    placeholder="Enter a title"
                    placeholderTextColor={"lightgrey"}                  
                    // onChangeText={handleTitleInput}
                  />
                </Block>

                <Text style={styles.secondaryHeader}>
                  Description
                </Text>
                <Block width={width * 0.8 - 20}  style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input}
                    placeholder="Provide post details here"
                    placeholderTextColor={"lightgrey"}
                    height={height * 0.3}
                    textAlignVertical={'top'}
                    paddingTop={10}
                    multiline={true}
                    // onChangeText={handleDescriptionInput}
                  />
                </Block>
                <ImagePicker></ImagePicker>
                <Pressable
                  style={[styles.imageButton]}
                  onPress={() => {uploadImageButton();}}>
                  <Text style={styles.textStyle}>UPLOAD IMAGE</Text>
                </Pressable>
                
                <Pressable
                  style={[styles.button, styles.buttonClose]}
                  onPress={() => {this.setState({modalVisible: !modalVisible}); handleCreatePostBtnClick();}}>
                  <Text style={styles.textStyle}>CREATE POST</Text>
                </Pressable>

              </View>
            </View>
          </View>
          </ScrollView>
        </Modal>
        {/* button that shows before opening modal */}
        <Pressable
          style={[styles.button, styles.buttonOpen]}
          onPress={() => this.setState({modalVisible: true})}>
          <Text style={styles.textStyle}>New Post</Text>
        </Pressable>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  exit: {
    margin:10,
    marginBottom: 0,
    justifyContent:"flex-end",
    width: 15,
    height: 15,
  },

  header: {
    fontSize: 25,
    fontWeight: 'bold',
    color: "#525F7F",
    marginBottom: 20,
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
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
   
  },
  modalView: {
    margin: 20,
    backgroundColor: 'white',
    paddingTop: 0,
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },
  imageButton: {
    borderRadius: 5,
    padding: 5,
    elevation: 2,
    // height:8,
    backgroundColor:'lightgrey',
 
  },

  // matt's added styles above ^^^

  centeredView: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 22,
  },

  modalViewInside: {
    padding: 25,
    paddingTop: 0,
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
    padding: 10,
    marginTop:10,
    marginBottom: 15,
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


export default NewPostModal;