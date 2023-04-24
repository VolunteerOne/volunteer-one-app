import { StyleSheet, Dimensions, Alert } from "react-native";
import { Block, Text, Button } from "galio-framework";
import { argonTheme } from "../../constants";
import { events } from "../../constants/HomeTab/event_details";
import { useState } from "react";
import LikeButton from "./LikeButton";
const { width } = Dimensions.get("screen");
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
  //button state
  const [isDisabled, setIsDisabled] = useState(false);
  //asks user to confirm signing up for an event, if they press yes, button is disabled
  const confirmationAlert = () =>
    Alert.alert("Sign up for event?", "", [
      {
        text: "Cancel",
        onPress: () => console.log("Cancel Pressed"),
        style: "cancel",
      },
      { text: "OK", onPress: () => handleSignUp() },
    ]);

  //if user signs up for event, then the button is disabled
  const handleSignUp = () => {
    setIsDisabled(true);
    Alert.alert("Organization has been notified!");
  };

  let eventDetails = events.find((o) => o.id === eventID);
  //if the eventID is found, display data
  if (eventDetails) {
    //retrieves and formats the body of the event post
    let bodyContent = eventDetails["eventBody"].map(function (data, index) {
      return (
        <Block key={index}>
          <Text style={styles.descriptionTitle}>{data["title"]}</Text>
          <Text>{data["description"]}</Text>
        </Block>
      );
    });
    //retrieves and formats the company info of the event post
    let companyInfo = eventDetails["companyInfo"].map(function (data, index) {
      return (
        <Block key={index}>
          <Text style={styles.descriptionTitle}>{data["title"]}</Text>
          <Text>{data["description"]}</Text>
        </Block>
      );
    });
    //Event Details gets returned
    return (
      <Block style={[styles.card, styles.shadowProp]}>
        <Block gap={8}>
          <Block row>
            <Text style={styles.headerTitle}>{eventDetails["title"]}</Text>
            <LikeButton />
          </Block>
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
          <Button
            disabled={isDisabled}
            color="primary"
            onPress={confirmationAlert}
            style={isDisabled ? styles.disabled : styles.signupButton}
          >
            <Text bold size={14} color={argonTheme.COLORS.WHITE}>
              {isDisabled ? "Sign up Successful" : "Sign Up"}
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
    gap: 8,
    marginBottom: 10,
  },
  signupButton: {
    width: width * 0.8,
    marginTop: 25,
  },
  disabled: {
    width: width * 0.8,
    marginTop: 25,
    opacity: 0.6,
  },
});

export default EventDetails;
