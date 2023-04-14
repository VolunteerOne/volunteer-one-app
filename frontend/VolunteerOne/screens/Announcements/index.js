import React from "react";
import { StyleSheet, Dimensions, ScrollView, Text } from "react-native";
import { Button } from "../../components";
import { Block, theme } from "galio-framework";
import EventCard from "../../components/EventCard";
import { following } from "../../constants/announcements_followingtab";
import { all } from "../../constants/announcements_alltab";
import argonTheme from "../../constants/Theme";
import NewAnnouncementModal from "../../components/Modals/NewAnnouncementModal";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

const { width } = Dimensions.get("screen");

/** ==================================== Announcements Tab ==================================== **/

class Announcements extends React.Component {
  renderArticles = () => {
    //route prop contains toggle parameter that tells the page to render content for the followers tab or all tab
    const { route } = this.props;
    //by default show followers page
    let toggle = true;
    //depending on if the user clicks on Followers button or All button, the data gets generated differently
    route.params ? (toggle = route.params.toggle) : (toggle = true);
    //list to display all the events
    var eventsList = [];
    if (toggle) {
      //display followers data
      eventsList = following.map(function (data) {
        return <EventCard key={data["id"]} data={data} />;
      });
    } else {
      //display all data
      eventsList = all.map(function (data) {
        return <EventCard key={data["id"]} data={data} />;
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

  state = {
    modalVisible: false,
  };

  render() {
    const { modalVisible } = this.state;

    const handleModalVisible = () => {
      this.setState({ modalVisible: !modalVisible });
    };

    return (
  
      <Block flex center style={styles.home}>
{/* ----------------------- new announcement button ----------------------- */}
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
            visible={this.state.modalVisible}
            setState={handleModalVisible}
          />
        )}

{/* ----------------------- render articles ----------------------- */}
        {this.renderArticles()}
      </Block>
    );
  }
}

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
