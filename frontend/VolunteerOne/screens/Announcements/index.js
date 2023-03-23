import React from "react";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import { Block, theme } from "galio-framework";
import EventCard from "../../components/EventCard";
import { dataList } from "../../constants/announcements";
const { width } = Dimensions.get("screen");

class Announcements extends React.Component {
  renderArticles = () => {
    var eventsList = dataList.map(function (data) {
      return <EventCard key={data["id"]} data={data} />;
    });

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
