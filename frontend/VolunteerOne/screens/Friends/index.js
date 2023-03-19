import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import { Button } from "../../components";

// import { Card } from '../../components';
import NotificationItem from '../../components/NotificationItem';
import ProfileItem from '../../components/ProfileItem';
// import articles from '../../constants/articles';
import notifications from '../../constants/notifications';
import profiles from '../../constants/profiles';

// <matt>
// import { Block, Button, Text, theme } from "galio-framework";
// </matt>

const { width } = Dimensions.get('screen');

class Home extends React.Component {
  renderNotifications = () => {
    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
 
        <Block flex>
          {profiles.map(profile => (
             <ProfileItem item={profile} horizontal />
          ))}
        </Block>
    
      </ScrollView>
    )
  }


  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderNotifications()}
      </Block>
    );
  }


}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default Home;
