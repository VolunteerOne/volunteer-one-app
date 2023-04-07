
import React from "react";
import {
  Alert,
  Modal,
  StyleSheet,
  Pressable,
  View,
  Dimensions,
  TextInput,
  Image,
  ScrollView,
} from "react-native";
import { Block, Text, theme } from "galio-framework";
import { argonTheme } from "../../constants";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import ImagePicker from './ImagePicker.js';

const { width, height } = Dimensions.get("screen");


/** ==================================== New Announcement Modal Component ==================================== **/

class NewEventModal extends React.Component {
  
  state = {
    user: "",
    datetime: new Date(),
    title: "",
    description: "",
  };

  render() {

    const handleAddNewClick = () => {
      console.log("Adding New Event ",this.state)
      // post to db
    }

    return (
      <View style={styles.centeredView}>
        <Modal
          animationType="fade"
          transparent={true}
          visible={this.props.visible}
          onRequestClose={() => {
            console.log("Modal has been closed.");
            this.props.setState();
          }}
        >
          
          <View style={[styles.centeredView, styles.modalViewOutside]}>
          <ScrollView contentContainerStyle={{ flexGrow: 1, justifyContent: 'center' }}> 
            <View style={styles.modalView}>
            
              {/* exit modal */}
              <Pressable
                onPress={() => this.props.setState()}
                style={{ alignItems: "flex-end", margin: 5 }}
              >
                <MaterialCommunityIcons
                  size={24}
                  name="close"
                  color={theme.COLORS.ICON}
                />
              </Pressable>

              <View style={styles.modalViewInside}>
                <Text style={styles.header}>New Event</Text>

                <Text style={styles.secondaryHeader}>Post title</Text>
                <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input}
                    placeholder="Enter a title"
                    placeholderTextColor={"lightgrey"}
                    onChangeText={(e) => this.setState({ title: e })}
                  />
                </Block>

                <Text style={styles.secondaryHeader}>Description</Text>
                <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input}
                    placeholder="Provide announcement details here"
                    placeholderTextColor={"lightgrey"}
                    height={height * 0.3}
                    textAlignVertical={"top"}
                    paddingTop={10}
                    multiline={true}
                    onChangeText={(e) => this.setState({ description: e })}
                  />
                </Block>
                <ImagePicker></ImagePicker>
                <Pressable
                  style={[styles.button, styles.buttonClose]}
                  onPress={() => {
                    this.props.setState(); 
                    handleAddNewClick();
                  }}
                >
                  <Text style={styles.textStyle}>CREATE EVENT</Text>
                </Pressable>
              </View>
              
            </View>
            </ScrollView>
          </View>
          
        </Modal>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  header: {
    fontSize: 25,
    fontWeight: "bold",
    color: "#525F7F",
    marginBottom: 20,
  },
  secondaryHeader: {
    fontSize: 17,
    fontWeight: "bold",
    color: "#525F7F",
    marginBottom: 5,
  },
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    borderWidth: 0.5,
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
    backgroundColor: "white",
    shadowColor: "#000",
    shadowOffset: {
      width: 5,
      height: 5,
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },

  // matt's added styles above ^^^

  centeredView: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    // marginTop: 22,
  },
  modalViewOutside: {
    backgroundColor: 'rgba(52, 52, 52, 0.75)',    // changed opacity of background when modal is open 
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
  buttonClose: {
    backgroundColor: "#5e72e4",
    padding: 10,
    marginTop: 10,
  },
  textStyle: {
    color: "white",
    fontWeight: "bold",
    textAlign: "center",
  },
});

export default NewEventModal;
