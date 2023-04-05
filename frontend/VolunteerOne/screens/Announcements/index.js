import React from "react";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import { Block, theme } from "galio-framework";
import EventCard from "../../components/EventCard";
import { following } from "../../constants/announcements_followingtab";
import { all } from "../../constants/announcements_alltab";
const { width } = Dimensions.get("screen");

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

  render() {
    return (
      <Block flex center style={styles.home}>
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
});

export default Announcements;
