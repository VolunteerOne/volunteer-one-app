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
  KeyboardAvoidingView,
} from "react-native";
import { Block, Text, theme } from "galio-framework";
import { argonTheme } from "../../constants";
import { Icon, Input } from "../../components";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
// import ImagePicker from "./ImagePicker.js";

const { width, height } = Dimensions.get("screen");

/** ==================================== New Skill Modal Component ==================================== **/

class AddNewSkillModal extends React.Component {
  state = {
    skill: "",
  };

  render() {
    const handleAddNewClick = () => {
      this.props.addSkill(this.state);
    };

    return (
      <View style={styles.centeredView}>
        <KeyboardAvoidingView style={{ flex: 1 }} behavior="padding" enabled>
          <Modal
            animationType="fade"
            transparent={true}
            visible={this.props.visible}
            onRequestClose={() => {
              this.props.setState();
            }}
          >
            <View style={[styles.centeredView, styles.modalViewOutside]}>
              <ScrollView
                contentContainerStyle={{
                  flexGrow: 1,
                  justifyContent: "center",
                }}
              >
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
                    <Text style={styles.header}>Add New Skill</Text>

                    <Block width={width * 0.8} style={{ marginBottom: 15 }}>
                      <Input
                        borderless
                        placeholder="New Skill"
                        iconContent={
                          <MaterialCommunityIcons
                            paddingLeft={5}
                            paddingRight={5}
                            size={16}
                            name="plus"
                            color={theme.COLORS.ICON}
                          />
                        }
                        onChangeText={(e) => this.setState({ skill: e })}
                      />
                    </Block>
                    <Pressable
                      style={[styles.button, styles.buttonClose]}
                      onPress={() => {
                        this.props.setState();
                        handleAddNewClick();
                      }}
                    >
                      <Text style={styles.textStyle}>ADD</Text>
                    </Pressable>
                  </View>
                </View>
              </ScrollView>
            </View>
          </Modal>
        </KeyboardAvoidingView>
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
  centeredView: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    // marginTop: 22,
  },
  modalViewOutside: {
    backgroundColor: "rgba(52, 52, 52, 0.75)", // changed opacity of background when modal is open
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

export default AddNewSkillModal;
