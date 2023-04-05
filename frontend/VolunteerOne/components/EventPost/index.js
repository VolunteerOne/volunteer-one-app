import { StyleSheet, Dimensions } from "react-native";
import { Block, Text, Button } from "galio-framework";
import { argonTheme } from "../../constants";
import { events } from "../../constants/event_details";
const { width, height } = Dimensions.get("screen");

const EventPost = ({ eventID }) => {
  if (eventID in events) {
    let eventDetails = events[eventID];
    //display all data
    let bodyContent = eventDetails["eventBody"].map(function (data) {
      return (
        <>
          <Text style={styles.descriptionTitle}>{data["title"]}</Text>
          <Text>{data["description"]}</Text>
        </>
      );
    });
    return (
      <Block style={[styles.card, styles.shadowProp]}>
        <Block style={styles.header}>
          <Text style={styles.headerTitle}>{eventDetails["title"]}</Text>
          <Text style={styles.headerText}>{eventDetails["organization"]}</Text>
          <Text style={styles.headerText}>{eventDetails["datePosted"]}</Text>
        </Block>
        <Block style={styles.divider}></Block>
        <Text style={[styles.headerText, styles.alignRight]}>
          {eventDetails["interestedPeople"]} people are interested
        </Text>
        <Block style={styles.body}>{bodyContent}</Block>
        <Block middle>
          <Button color="primary" style={styles.signupButton}>
            <Text bold size={14} color={argonTheme.COLORS.WHITE}>
              Sign up
            </Text>
          </Button>
        </Block>
      </Block>
    );
  }
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#FFFFFF",
    width: "100%",
    minWidth: "100%",
    borderRadius: 10,
    padding: 15,
    margin: 10,
  },
  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
  headerTitle: {
    fontSize: 20,
    fontWeight: "bold",
    color: "#32325D",
  },
  headerText: {
    color: "rgba(50, 50, 93, 0.5)",
  },
  header: {
    flexDirection: "column",
    justifyContent: "flex-start",
    gap: 8,
  },
  descriptionTitle: {
    color: "#5E72E4",
    fontWeight: 700,
  },
  divider: {
    borderBottomColor: "#E9ECEF",
    borderBottomWidth: 1,
    marginVertical: 10,
  },
  alignRight: {
    textAlign: "right",
  },
  body: {
    flexDirection: "column",
    justifyContent: "flex-start",
    gap: 8,
    marginBottom: 10,
  },
  signupButton: {
    width: width * 0.8,
    marginTop: 25,
  },
});

export default EventPost;
