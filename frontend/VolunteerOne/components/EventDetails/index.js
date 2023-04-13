import { StyleSheet, Dimensions } from "react-native";
import { Block, Text, Button } from "galio-framework";
import { argonTheme } from "../../constants";
import { events } from "../../constants/event_details";
const { width, height } = Dimensions.get("screen");
/*
Description:
  This component displays the event details of an organization's post when a user presses
  "Click to View Event". If the data for an event can't be found, it will display 
  "No event found"

Props received:
  eventID - when a user clicks on "View event", the event ID gets passed to the 
            EventDetails Component. For now the EventDetails component will look 
            for the corresponding event in a file called event_details.js in the 
            constants folder. 
*/
const EventDetails = ({ eventID }) => {
  //if the eventID is found, display data
  if (eventID in events) {
    let eventDetails = events[eventID];
    //retrieves and formats the body of the event post
    let bodyContent = eventDetails["eventBody"].map(function (data) {
      return (
        <>
          <Text style={styles.descriptionTitle}>{data["title"]}</Text>
          <Text>{data["description"]}</Text>
        </>
      );
    });
    //retrieves and formats the company info of the event post
    let companyInfo = eventDetails["companyInfo"].map(function (data) {
      return (
        <>
          <Text style={styles.descriptionTitle}>{data["title"]}</Text>
          <Text>{data["description"]}</Text>
        </>
      );
    });
    //Event Details gets returned
    return (
      <Block style={[styles.card, styles.shadowProp]}>
        <Block style={styles.headerContent}>
          <Text style={styles.headerTitle}>{eventDetails["title"]}</Text>
          <Text style={styles.headerText}>{eventDetails["organization"]}</Text>
          <Text style={styles.headerText}>{eventDetails["datePosted"]}</Text>
        </Block>
        <Block style={styles.divider}></Block>
        <Text style={[styles.headerText, styles.alignRight]}>
          {eventDetails["interestedPeople"]} people are interested
        </Text>
        <Block style={styles.body}>{bodyContent}</Block>
        <Block style={styles.divider}></Block>
        <Block style={styles.body}>{companyInfo}</Block>
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
  //if no event is found, display error message
  else {
    return <Text>No event found</Text>;
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
    fontSize: 15,
    fontWeight: 500,
  },
  headerContent: {
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
    marginVertical: 15,
  },
  alignRight: {
    textAlign: "right",
  },
  body: {
    marginTop: 10,
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

export default EventDetails;