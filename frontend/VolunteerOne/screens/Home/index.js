import React from "react";
import { StyleSheet, Dimensions, ScrollView } from "react-native";
import { Block, theme } from "galio-framework";
import EventCard from "../../components/EventCard";
const { width } = Dimensions.get("screen");

const dataList = [
  {
    id: 0,
    organization: "American Red Cross",
    name: "New Event",
    subject: "Help pack disaster relief bags",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
    date: "",
    location: "",
    timePosted: new Date("2023-03-01T00:00:00"),
    type: "event",
  },
  {
    id: 1,
    organization: "Habitat for Humanity",
    name: "New Event",
    subject: "Build and improve homes",
    description:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
    date: "",
    location: "",
    timePosted: new Date("2023-02-02T00:00:00"),
    type: "event",
  },
  {
    id: 2,
    organization: "Global Volunteers",
    announcement:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et enim ad minim veniam, quis nostrud exercitation ullamco laboris ",
    type: "announcment",
    timePosted: new Date("2021-03-01T00:00:00"),
  },
];

class Home extends React.Component {
  renderArticles = () => {
    const { navigation } = this.props;
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

export default Home;
