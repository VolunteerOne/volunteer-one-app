import React, { useEffect } from "react";
import { useState } from "react";
import {
  StyleSheet,
  Dimensions,
  Image,
  ScrollView,
  TouchableOpacity,
  View,
  Pressable,
} from "react-native";
import { Block, theme, Text } from "galio-framework";
const { width, height } = Dimensions.get("screen");

import { Images, argonTheme } from "../../constants";
import { TextInput } from "react-native-gesture-handler";
import mockData from "../../constants/ProfileTab/profile";
import { Button } from "../../components";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import AddNewSkillModal from "../../components/Modals/AddNewSkillModal";

// class Settings extends React.Component {
const Settings = ({ navigation }) => {
  // render() {

  const JESSICA = "Jessica Jones";
  const [user, setProfile] = useState(mockData[JESSICA]);
  const [userName, setUserName] = useState(user.name);
  const [userLocation, setUserLocation] = useState();
  const [userBio, setUserBio] = useState();
  const [userSkills, setUserSkills] = useState([
    "Crafts",
    "Drawing",
    "Handiwork",
    "Computer",
    "Painting",
  ]);
  const [modalVisible, setModalVisible] = useState(false);

  const handleAddNewSkill = (skillObj) => {
    const { skill } = skillObj;
    if (skill) setUserSkills([skill, ...userSkills]);
  };

  console.log(userSkills);

  useEffect(() => {
    setProfile(mockData[JESSICA]);
    setUserName(mockData[JESSICA].name);
    setUserLocation(`${mockData[JESSICA].city}, ${mockData[JESSICA].country}`);
    setUserBio(mockData[JESSICA].description);
  }, []);

  return (
    <ScrollView showsVerticalScrollIndicator={false}>
      <Block flex center style={styles.home} paddingTop={100} gap={15}>
        <Block middle style={styles.avatarContainer}>
          <Image source={{ uri: user.image }} style={styles.avatar} />
        </Block>
        <Block>
          <TouchableOpacity>
            <Text
              color={argonTheme.COLORS.ACTIVE}
              size={15}
              paddingBottom={25}
              style={{
                fontWeight: "bold",
              }}
            >
              Edit Image
            </Text>
          </TouchableOpacity>
        </Block>
        <Block style={styles.settingContainer}>
          <Block style={styles.inputKey}>
            <Text paddingLeft={10} color="#000000" size={14}>
              Name
            </Text>
          </Block>
          <Block>
            <TextInput
              width="100%"
              paddingLeft={50}
              color="#000000"
              style={styles.input}
              placeholder={userName}
              placeholderTextColor="#000000"
              onChangeText={setUserName}
            />
          </Block>
        </Block>
        <Block style={styles.settingContainer}>
          <Block style={styles.inputKey}>
            <Text paddingLeft={10} color="#000000" size={14}>
              Location
            </Text>
          </Block>
          <Block>
            <TextInput
              width="100%"
              paddingLeft={50}
              color="#000000"
              style={styles.input}
              placeholder={userLocation}
              placeholderTextColor="#000000"
              onChangeText={setUserLocation}
            />
          </Block>
        </Block>

        <View style={styles.bioContainer}>
          <View>
            <Text padding={10} color="#000000" size={14}>
              Bio
            </Text>
          </View>
          <View>
            <Text paddingLeft={10} paddingRight={10} color="#000000" size={16}>
              {userBio}
            </Text>
          </View>
        </View>

        <View style={[styles.bioContainer]}>
          <View style={{ flexDirection: "row", justifyContent: "flex-end" }}>
            <Text padding={10} color="#000000" size={14}>
              Skills
            </Text>
            <Pressable
              onPress={() => setModalVisible(true)}
              // style={{ alignItems: "flex-end", margin: 5 }}
            >
              <MaterialCommunityIcons
                // paddingLeft={5}
                padding={10}
                size={16}
                name="pencil-plus"
                color={theme.COLORS.ICON}
              />
            </Pressable>
          </View>
          <View
            style={{
              // padding: 5,
              width: "100%",
              flexDirection: "row",
              flexWrap: "wrap",
              justifyContent: "center",
            }}
          >
            {userSkills.map((skill) => (
              <Button
                key={skill}
                small
                style={{ backgroundColor: argonTheme.COLORS.DEFAULT }}
              >
                <Text size={12} color="white">
                  {skill}
                </Text>
              </Button>
            ))}
          </View>
          {modalVisible && (
            <AddNewSkillModal
              visible={modalVisible}
              setState={setModalVisible}
              addSkill={handleAddNewSkill}
            />
          )}
        </View>
      </Block>
    </ScrollView>
  );
  // }
};

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  skillBox: {
    backgroundColor: argonTheme.COLORS.PRIMARY,
    color: "white",
    borderRadius: 5,
    margin: 2,
    // size: 16,
    // alignItems: 'center',
    // justifyContent: 'center',
  },
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
  },
  inputKey: {
    width: width * 0.2,
  },
  settingContainer: {
    width: width * 0.8,
    height: height * 0.05,
    backgroundColor: "#FFFFFF",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    flex: 1,
    justifyContent: "left",
    paddingLeft: 5,
    shadowRadius: 8,
    shadowOpacity: 0.1,
    elevation: 1,
    flexDirection: "row",
    alignItems: "flex-start",
    alignItems: "center",
  },
  bioContainer: {
    width: width * 0.8,
    // height: height * 0.3,
    backgroundColor: "#FFFFFF",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    flex: 1,
    // paddingTop: 10,
    // paddingLeft: 5,
    padding: 10,
    shadowRadius: 8,
    shadowOpacity: 0.1,
    alignItems: "flex-start",
  },
  avatarContainer: {
    position: "relative",
    marginTop: -80,
  },
  avatar: {
    width: 124,
    height: 124,
    borderRadius: 62,
    borderWidth: 0,
  },
});

export default Settings;
