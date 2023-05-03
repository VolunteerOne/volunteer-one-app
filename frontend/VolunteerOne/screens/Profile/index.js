import React, { useState } from "react";
import {
  Alert,
  StyleSheet,
  Dimensions,
  ScrollView,
  Image,
  ImageBackground,
  Platform,
  TouchableOpacity,
  View,
} from "react-native";
import { Block, Text, theme } from "galio-framework";
import { Button } from "../../components";
import { Images, argonTheme } from "../../constants";
import { HeaderHeight } from "../../constants/utils";
import RecentActivityCard from "../../components/RecentActivityCard";
import UpcomingEventsCard from "../../components/UpcomingEventsCard";
import { useRoute } from "@react-navigation/native";

// constants
import mockData from "../../constants/ProfileTab/profile";

const { width, height } = Dimensions.get("screen");

const thumbMeasure = (width - 48 - 32) / 3;

/** ==================================== Profile Screen ==================================== **/

const ProfileScreen = ({ route, navigation }) => {
  // pass user name param when navigating to a profile - matt
  let anotherUser = false; // false if viewing own profile
  let theUser = "Jessica Jones"; // default user
  const [userSkills, setUserSkills] = useState(typeof route.params === "undefined" ? mockData[theUser].skills : mockData[route.params.theUser].skills);

  if (typeof route.params !== "undefined") {
    theUser = route.params.theUser;
    anotherUser = true; // viewing a different user
  }
  console.log("Viewing user: ", theUser);

  let isVolunteer = true;
  if (mockData[theUser].userType == "organization") {
    isVolunteer = false;
  }

  const [followText, setFollowText] = useState("FOLLOW");

  const handleConnectBtnPress = () => {
    console.log("connect btn pressed");
  };
  const handleMessageBtnPress = () => {
    console.log("Message btn pressed");
  };

  const handleFollowBtnPress = () => {
    setFollowText("UNFOLLOW");
    Alert.alert("Followed!");
  };

  // const handleViewAllRecentActivityBtn = () => {
  //   console.log("handleViewAllRecentActivityBtn");
  // };

  // const handleViewAllUpcomingEventsBtn = () => {
  //   console.log("handleViewAllUpcomingEventsBtn");
  // };

  return (
    <ScrollView>
      <Block flex style={styles.profile}>
        <Block flex>
          <ImageBackground
            source={Images.ProfileBackground}
            style={styles.profileContainer}
            imageStyle={styles.profileBackground}
          >
            <ScrollView
              showsVerticalScrollIndicator={false}
              style={{ width, marginTop: "25%" }}
            >
              <Block flex style={styles.profileCard}>
                <Block middle style={styles.avatarContainer}>
                  <Image
                    source={{ uri: mockData[theUser].image }}
                    style={styles.avatar}
                  />
                </Block>
                <Block style={styles.info}>
                  <Block
                    middle
                    row
                    space="evenly"
                    style={{ marginTop: 20, paddingBottom: 24 }}
                  >
                    {isVolunteer ? (
                      <Button
                        small
                        style={{ backgroundColor: argonTheme.COLORS.INFO }}
                        onPress={handleConnectBtnPress}
                      >
                        CONNECT
                      </Button>
                    ) : null}

                    {isVolunteer ? (
                      <Button
                        small
                        style={{ backgroundColor: argonTheme.COLORS.DEFAULT }}
                        onPress={handleMessageBtnPress}
                      >
                        {"MESSAGE"}
                      </Button>
                    ) : (
                      <Button
                        small
                        style={{ backgroundColor: argonTheme.COLORS.DEFAULT }}
                        onPress={() => handleFollowBtnPress()}
                      >
                        {followText}
                      </Button>
                    )}
                  </Block>
                  {isVolunteer && (
                    <Block
                      row
                      flex
                      space="between"
                      style={{ marginBottom: 35 }}
                    >
                      <Block middle>
                        <Text
                          bold
                          size={18}
                          color="#525F7F"
                          style={{ marginBottom: 4 }}
                        >
                          {mockData[theUser].hours}
                        </Text>
                        <Text size={12} color={argonTheme.COLORS.TEXT}>
                          Hours
                        </Text>
                      </Block>
                      <TouchableOpacity
                        onPress={() =>
                          navigation.navigate("ViewFriends", {
                            theUser: theUser,
                          })
                        }
                      >
                        <Block middle>
                          <Text
                            bold
                            color="#525F7F"
                            size={18}
                            style={{ marginBottom: 4 }}
                          >
                            {mockData[theUser].friends.value}
                          </Text>
                          <Text size={12} color={argonTheme.COLORS.TEXT}>
                            Friends
                          </Text>
                        </Block>
                      </TouchableOpacity>
                      <TouchableOpacity
                        onPress={() =>
                          navigation.navigate("ViewFollowing", {
                            theUser: theUser,
                          })
                        }
                      >
                        <Block middle>
                          <Text
                            bold
                            color="#525F7F"
                            size={18}
                            style={{ marginBottom: 4 }}
                          >
                            {mockData[theUser].following.value}
                          </Text>
                          <Text size={12} color={argonTheme.COLORS.TEXT}>
                            Following
                          </Text>
                        </Block>
                      </TouchableOpacity>
                    </Block>
                  )}
                </Block>
                <Block flex>
                  <Block middle style={styles.nameInfo}>
                    <Text bold size={28} color="#32325D">
                      {mockData[theUser].name}
                      {isVolunteer && `, ${mockData[theUser].age}`}
                    </Text>
                    <Text size={16} color="#32325D" style={{ marginTop: 10 }}>
                      {mockData[theUser].city}, {mockData[theUser].country}
                    </Text>
                  </Block>
                  <Block middle style={{ marginTop: 30, marginBottom: 16 }}>
                    <Block style={styles.divider} />
                  </Block>
                  <Block middle style={{ marginBottom: 5 }}>
                    {!isVolunteer && (
                      <Text
                        bold
                        size={14}
                        color={argonTheme.COLORS.PRIMARY}
                        style={{ marginBottom: 10 }}
                      >
                        Mission Statement
                      </Text>
                    )}
                    <Text
                      size={14}
                      color="#525F7F"
                      style={{ textAlign: "center" }}
                    >
                      {mockData[theUser].description}
                    </Text>
                  </Block>
                  {!isVolunteer && (
                    <>
                      <Block row space="between">
                        <Text
                          bold
                          size={16}
                          color="#525F7F"
                          style={{ marginTop: 12 }}
                        >
                          Upcoming Events
                        </Text>
                        {/* <Button
                          small
                          color="transparent"
                          textStyle={{
                            color: "#5E72E4",
                            fontSize: 12,
                            marginLeft: 24,
                          }}
                          onPress={handleViewAllUpcomingEventsBtn}
                        >
                          View all
                        </Button> */}
                      </Block>
                      <Block style={{ paddingBottom: -HeaderHeight * 2 }}>
                        <Block row space="between" style={{ flexWrap: "wrap" }}>
                          {mockData[theUser].upcomingEvents.map((post) => (
                            <UpcomingEventsCard data={post} />
                          ))}
                        </Block>
                      </Block>
                    </>
                  )}

                  { isVolunteer &&
                    <Block space="between" style={{ marginTop: 12 }}>
                    <Text
                      bold
                      size={16}
                      color="#525F7F"
                      style={{ marginTop: 12 }}
                    >
                      Skills
                    </Text>
                    <View
                      style={{
                        // padding: 5,
                        width: "100%",
                        flexDirection: "row",
                        flexWrap: "wrap",
                        // justifyContent: "center",
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
                  </Block>
                  }

                  <Block row space="between" style={{ marginTop: 12 }}>
                    <Text
                      bold
                      size={16}
                      color="#525F7F"
                      style={{ marginTop: 12 }}
                    >
                      Recent Activity
                    </Text>
                    {/* <Button
                      small
                      color="transparent"
                      textStyle={{
                        color: "#5E72E4",
                        fontSize: 12,
                        marginLeft: 24,
                      }}
                      onPress={handleViewAllRecentActivityBtn}
                    >
                      View all
                    </Button> */}
                  </Block>
                  <Block style={{ paddingBottom: -HeaderHeight * 2 }}>
                    <Block row space="between" style={{ flexWrap: "wrap" }}>
                      {mockData[theUser].recentActivity.map((post, i) => (
                        <RecentActivityCard key={i} data={post} />
                      ))}
                    </Block>
                  </Block>
                </Block>
              </Block>
            </ScrollView>
          </ImageBackground>
        </Block>
      </Block>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  profile: {
    marginTop: Platform.OS === "android" ? -HeaderHeight : 0,
    // marginBottom: -HeaderHeight * 2,
    flex: 1,
  },
  profileContainer: {
    width: width,
    height: height * 0.9, // needs to account for bottom tab navigator
    padding: 0,
    zIndex: 1,
  },
  profileBackground: {
    width: width,
    height: height / 2,
  },
  profileCard: {
    // position: "relative",
    padding: theme.SIZES.BASE,
    marginHorizontal: theme.SIZES.BASE,
    marginTop: 65,
    borderTopLeftRadius: 6,
    borderTopRightRadius: 6,
    backgroundColor: theme.COLORS.WHITE,
    shadowColor: "black",
    shadowOffset: { width: 0, height: 0 },
    shadowRadius: 8,
    shadowOpacity: 0.2,
    zIndex: 2,
  },
  info: {
    paddingHorizontal: 40, // this makes the friends count not centered
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
  divider: {
    width: "90%",
    borderWidth: 1,
    borderColor: "#E9ECEF",
  },
  thumb: {
    borderRadius: 4,
    marginVertical: 4,
    alignSelf: "center",
    width: thumbMeasure,
    height: thumbMeasure,
  },
});

export default ProfileScreen;
