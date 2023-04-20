import React, { useState } from "react";
import { StyleSheet, Dimensions, ScrollView, Text } from "react-native";
import { Button } from "../../components";
import { Block, theme } from "galio-framework";
import EventCard from "../../components/EventCard";
import { following } from "../../constants/HomeTab/announcements_followingtab";
import { all } from "../../constants/HomeTab/announcements_alltab";
import profiles from "../../constants/ProfileTab/profile";
import argonTheme from "../../constants/Theme";
import NewAnnouncementModal from "../../components/Modals/NewAnnouncementModal";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

const { width } = Dimensions.get("screen");

/** ==================================== Announcements Tab ==================================== **/

const Announcements = ({ navigation, route }) => {

  const JESSICA = "Jessica Jones";
  const [user, setUser] = useState(profiles[JESSICA]);

  const [modalVisible, setModalVisible] = useState(false);
  const [announcements, setAnnouncements] = useState(following)

  const handleModalVisible = () => {
    setModalVisible(!modalVisible);
  };

  const addNewAnnouncement = (data) => {
    // console.log('Adding New Announcement', data)
    const newAnnouncement = {
      id: announcements.length +1,
      organization: {
        name: user.name,
        image: user.image
      },
      announcement: `${data['title']}\n\n${data['description']}`,
      type: "announcement",
      timePosted: data['datetime'],
    }
    // console.log(newAnnouncement)
    setAnnouncements([newAnnouncement, ...announcements])
  }

  const renderArticles = () => {
    //by default show followers page
    let toggle = true;
    route.params ? (toggle = route.params.toggle) : (toggle = true);
    //list to display all the events
    var eventsList = [];
    if (toggle) {
      //display followers data
      eventsList = announcements.map((data, i) => {
        return <EventCard key={i} data={data} />;
      });
    } else {
      // display all data
      eventsList = all.map((data) => {
        return <EventCard key={i} data={data} />;
      });
    }

    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.articles}
      >
        <Block flex center>
          {eventsList}
        </Block>
      </ScrollView>
    );
  };

  return (
    <Block flex center style={styles.home}>
      <Block middle>
        <Button
          color="primary"
          style={styles.button}
          onPress={() => handleModalVisible()}
        >
          <Block row middle>
            <MaterialCommunityIcons
              size={24}
              name="plus-box-outline"
              color={theme.COLORS.WHITE}
            />
            <Text bold size={14} style={styles.buttonTitle}>
              New Announcement
            </Text>
          </Block>
        </Button>
      </Block>
      {modalVisible && (
        <NewAnnouncementModal
          visible={modalVisible}
          handleModalVisible={handleModalVisible}
          addNewAnnouncement={addNewAnnouncement}
        />
      )}

      {renderArticles()}
    </Block>
  );
};

const styles = StyleSheet.create({
  home: {
    width: width,
  },
  articles: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
  button: {
    marginTop: theme.SIZES.BASE,
    marginBottom: 0,
    width: width * 0.9,
  },
  buttonTitle: {
    paddingLeft: 5,
    lineHeight: 19,
    fontWeight: "600",
    color: argonTheme.COLORS.WHITE,
  },
});

export default Announcements;
